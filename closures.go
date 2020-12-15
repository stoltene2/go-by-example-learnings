package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	i := intSeq()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}
