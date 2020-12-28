package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	done := make(chan bool, 1)

	go func() {
		for elem := range queue {
			fmt.Println(elem)
		}

		done <- true
	}()

	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	close(queue)
	<-done
}
