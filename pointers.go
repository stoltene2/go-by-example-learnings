package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(pval *int) {
	*pval = 0
}

func main() {
	ival := 1
	fmt.Println("ival:", ival)
	zeroval(ival)
	fmt.Println("zeroval(ival):", ival)

	iptr := 1
	fmt.Println("iptr:", iptr)
	zeroptr(&iptr)
	fmt.Println("zeroval(iptr):", iptr, &iptr)

}
