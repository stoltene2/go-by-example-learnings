package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://goexample.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
}
