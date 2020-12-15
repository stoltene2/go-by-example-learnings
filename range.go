package main

import (
	"fmt"
)

func main() {
	nums := []int{2,3,4}
	sum := 0
	fmt.Println("sum:", sum)

	for _, n := range nums {
		sum += n
	}

	fmt.Println("sum:", sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println("index:", i)
		}
	}


	m := map[string]string{"a": "apple", "b": "bannana"}

	for k, _ := range m {
		fmt.Printf("%s -> %s\n", k, m[k])
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}


}
