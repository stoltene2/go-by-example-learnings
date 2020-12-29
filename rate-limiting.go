package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   This can also be used for sampling some data source
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200*time.Millisecond)

	for req := range requests {
		<-limiter // wait 200ms before servicing the queue
		fmt.Println("requests", req, time.Now())
	}

	//////////////////////////////

	// Max channel size is 3, so we'll make 3 slots available
	// right away
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// Every 200ms add a new slot for processing
		for t := range time.Tick(200*time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
        <-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
