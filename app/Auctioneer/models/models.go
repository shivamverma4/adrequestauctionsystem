package auctioneermodels

import "sync"

type AuctionStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AppState struct {
	sync.Mutex
	AuctionList []AuctionStruct `json:"auction_list"`
	LiveAuction AuctionStruct   `json:"live_auction"`
}
