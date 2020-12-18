package bidder

import (
	biddermodels "adrequestauctionsystem/app/bidder/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type bidderHandler func(time.Duration, int) biddermodels.RequestHandlerFunction

func StartBidderServer(bidder biddermodels.BidderStruct, handlerFunc bidderHandler) error {
	m := http.NewServeMux()
	s := http.Server{Addr: ":" + strconv.Itoa(bidder.Port), Handler: m}

	const auctionNotifyURL = "/auction/notification"
	m.HandleFunc(auctionNotifyURL, handlerFunc(bidder.Delay, bidder.Id))
	if err := s.ListenAndServe(); err != nil {
		log.Print("ListenAndServe: ", err)
		return err
	}
	fmt.Println("Server started at " + s.Addr)
	return nil
}
