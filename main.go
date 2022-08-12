package main

import (
	"bbrombacher/cryptoautotrader/coinbase"
	"bbrombacher/cryptoautotrader/transforms"
	"fmt"
	"log"
)

var socketUrl = "wss://ws-feed.pro.coinbase.com"

func main() {

	coinbaseClient := coinbase.New(socketUrl)

	feedParams := transforms.MakeStartTickerParams([]string{"ETH-USD"})
	tickerID, err := coinbaseClient.StartTickerFeed(feedParams)
	if err != nil {
		log.Fatal("failed to start ticker", err.Error())
	}

	for {
		msg, err := coinbaseClient.GetTickerMessages(tickerID)
		if err != nil {
			log.Fatal("could not get ticker message", err.Error())
		}
		fmt.Printf("msg %#v\n", msg)
	}
}
