package auctioneer

import (
	"adrequestauctionsystem/app/utils"
	"fmt"
	"time"
	"net/http"
	auctioneermodels "adrequestauctionsystem/app/auctioneer/models"
	biddermodels "adrequestauctionsystem/app/bidder/models"
	bidderController "adrequestauctionsystem/app/bidder"

	"github.com/labstack/echo"
)

var (
	allAuctions auctioneermodels.AppState
	allottedIds map[int]struct{}
)

func init() {
	allAuctions = auctioneermodels.AppState{}
	allottedIds = map[int]struct{}{}
}

func HandleCreateAuction(c echo.Context) (err error) {

	auctionData := new(auctioneermodels.AuctionStruct)
	if err = c.Bind(auctionData); err != nil {
		return
	}

	newAuctionId := utils.GetUniqueAllotedId(allottedIds)
	if newAuctionId == 0 {
		resp := utils.CustomHTTPResponse{}
		resp.Data = 0
		resp.Message = "Unique Auction ID cannot be generated"
		return c.JSON(http.StatusBadRequest, resp)
	}
	auctionData.Id = newAuctionId

	allAuctions.Lock()
	defer allAuctions.Unlock()

	allAuctions.AuctionList = append(allAuctions.AuctionList, *auctionData)
	fmt.Println("auctionData: ", auctionData)
	fmt.Println("allAuctionList: ", allAuctions)
	resp := utils.CustomHTTPResponse{}
	resp.Data = auctionData
	resp.Message = "Auction created"
	return c.JSON(http.StatusOK, resp)
}

func HandleGetAllAuctionsList(c echo.Context) (err error) {
	
	resp := utils.CustomHTTPResponse{}
	if len(allAuctions.AuctionList) == 0 {
		resp.Data = nil
		resp.Message = "No auctions exists"
	} else {
		resp.Data = allAuctions.AuctionList
		resp.Message = "All auctions fetched"
	}
	return c.JSON(http.StatusOK, resp)
}

func HandleBidRound(c echo.Context) (err error) {
	auctionLength := len(allAuctions.AuctionList)

	if auctionLength == 0 {
		resp := utils.CustomHTTPResponse{}
		resp.Data = nil
		resp.Message = "No Auction exists, create an auction to conduct bid round"
		return c.JSON(http.StatusBadRequest, resp)
	}
	allAuctions.LiveAuction, allAuctions.AuctionList = allAuctions.AuctionList[0], allAuctions.AuctionList[1:]

	bidEntriesChannel, biddersObj := make(chan biddermodels.BidResponse, 10), bidderController.GetAllBiddersList()
	select {
		case bid := <-bidEntriesChannel:
			fmt.Println("received bid", bid)

		default:
			fmt.Println("no bids received")
	}

	go SendAuctionNotification(biddersObj, bidEntriesChannel)

	timer := time.NewTimer(200 * time.Millisecond)
	<-timer.C
	close(bidEntriesChannel)

	if len(bidEntriesChannel) == 0 {
		allAuctions.AuctionList = append(allAuctions.AuctionList, allAuctions.LiveAuction)
		resp := utils.CustomHTTPResponse{}
		resp.Data = map[string]interface{}{
			"auctionID": allAuctions.LiveAuction.Id,
			"bidder": nil,
		}
		resp.Message = "No bids received"
		return c.JSON(http.StatusBadRequest, resp)
	}

	maxBidder := biddermodels.BidResponse{}
	for oneBidder := range bidEntriesChannel {
		if oneBidder.Price > maxBidder.Price {
			maxBidder.BidderId = oneBidder.BidderId
			maxBidder.Price = oneBidder.Price
		}
	}

	resp := utils.CustomHTTPResponse{}
	resp.Data = map[string]interface{}{
		"auctionID": allAuctions.LiveAuction.Id,
		"bidder": maxBidder,
	}
	resp.Message = "Winning bidder for the auction fetched"
	return c.JSON(http.StatusOK, resp)
}
