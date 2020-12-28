package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500*time.Millisecond)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				fmt.Println("ticked at", t)
			}
		}
	}()

	time.Sleep(2*time.Second)
	ticker.Stop()
	done <- true // Tell go routine to stop running
	fmt.Println("stopped")
}
