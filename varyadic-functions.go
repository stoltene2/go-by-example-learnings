package main

import (
	"fmt"
)

func sum(nums ...int) int {
	t := 0
	for _, n := range nums {
        t += n
	}
	return t
}

func main() {
	fmt.Println("sum:", sum(1,2,3,4,5))

	nums := []int{1,2,3,4,5}
	fmt.Println("sum:", sum(nums...))
}
