package auctioneer

import (
	biddermodels "adrequestauctionsystem/app/bidder/models"
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

func SendAuctionNotification(bidder biddermodels.AppState, bidEntriesChannel chan biddermodels.BidResponse) {
	const baseURL, auctionNotifyURL = "http://127.0.0.1:", "/auction/notification"
	client := http.Client{Timeout: 200 * time.Millisecond}
	for _, v := range bidder.BidderList {
		url := baseURL + strconv.Itoa(v.Port) + auctionNotifyURL
		go sendRequest(client, url, bidEntriesChannel)
	}
}

func sendRequest(client http.Client, url string, channel chan biddermodels.BidResponse) biddermodels.BidResponse {
	var bidResponse biddermodels.BidResponse
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Response error: ", err)
		return bidResponse
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&bidResponse)
	select {
		case channel <- bidResponse:
			fmt.Println("post bid into the auction channel", bidResponse)
		default:
			fmt.Println("no bid posted")
	}
	return bidResponse
}
