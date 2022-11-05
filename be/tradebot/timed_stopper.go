package tradebot

import (
	"log"
	"time"
)

func Stop(duration int, stopFunc func()) {
	log.Printf("stop function initialized for %d seconds", duration)

	// make a timer
	stopTimer := time.NewTimer(time.Duration(duration) * time.Second)

	// when timer stops run provided func
	<-stopTimer.C
	stopFunc()
}
