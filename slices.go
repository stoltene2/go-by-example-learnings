package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
	c[0] = "aaaaaa"

	fmt.Println("s:", s)
	fmt.Println("copy: ", c)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
