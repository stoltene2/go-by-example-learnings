package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2*time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	t2 := time.NewTimer(time.Second)

	go func() {
		<-t2.C
		fmt.Println("Timer 2 fired")
	}()

	stopped := t2.Stop()

	if stopped {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2*time.Second)
}
