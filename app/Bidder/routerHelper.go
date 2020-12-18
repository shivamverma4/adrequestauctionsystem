package bidder

import (
	"adrequestauctionsystem/app/utils"
	"net/http"
	"time"
	"encoding/json"
	biddermodels "adrequestauctionsystem/app/bidder/models"

	"github.com/labstack/echo"
)

var (
	allBidders  biddermodels.AppState
	allottedIds map[int]struct{}
)

func init() {
	allBidders = biddermodels.AppState{}
	allottedIds = map[int]struct{}{}
}

func HandleCreateBidder(c echo.Context) (err error) {

	bidderData := new(biddermodels.BidderStruct)
	if err = c.Bind(bidderData); err != nil {
		return
	}

	if len(allBidders.BidderList) >= 100 {
		resp := utils.CustomHTTPResponse{}
		resp.Data = 0
		resp.Message = "Cannot create bidder, maximum number of bidders count reached"
		return c.JSON(http.StatusNotAcceptable, resp)
	}

	newBidderId := utils.GetUniqueAllotedId(allottedIds)
	if newBidderId == 0 {
		resp := utils.CustomHTTPResponse{}
		resp.Data = 0
		resp.Message = "Unique Auction ID cannot be generated"
		return c.JSON(http.StatusBadRequest, resp)
	}
	bidderData.Id = newBidderId

	if !utils.IsPortFree(bidderData.Port) {
		resp := utils.CustomHTTPResponse{}
		resp.Data = 0
		resp.Message = "This port is in use, try bidder with different port number"
		return c.JSON(http.StatusBadRequest, resp)
		return
	}

	go StartBidderServer(*bidderData, BidderNotificationHandler)

	allBidders.Lock()
	defer allBidders.Unlock()

	allBidders.BidderList = append(allBidders.BidderList, *bidderData)
	resp := utils.CustomHTTPResponse{}
	resp.Data = bidderData
	resp.Message = "Bidder created and bidder server is up"
	return c.JSON(http.StatusOK, resp)
}

func HandleGetAllBiddersList(c echo.Context) (err error) {
	
	resp := utils.CustomHTTPResponse{}
	if len(allBidders.BidderList) == 0 {
		resp.Data = nil
		resp.Message = "No bidders exists"
	} else {
		resp.Data = allBidders.BidderList
		resp.Message = "All bidders fetched"
	}
	return c.JSON(http.StatusOK, resp)
}

func BidderNotificationHandler(t time.Duration, id int) biddermodels.RequestHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t * time.Millisecond)

		var bidResp biddermodels.BidResponse
		bidResp.BidderId = id
		bidResp.Price = utils.GetRandomFloat()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bidResp)
	}
}

func GetAllBiddersList() biddermodels.AppState {
	return allBidders
}
