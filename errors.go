package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg+3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg+3, nil
}

func main() {

	for _, item := range []int{7,42} {
        if res, err := f1(item); err == nil {
			fmt.Println("item:",item, res)
		} else {
			fmt.Println("FAIL:", err)
		}
	}

	for _, item := range []int{7,42} {
        if res, err := f2(item); err == nil {
			fmt.Println("item:",item, res)
		} else {
			fmt.Println("FAIL:", err)
		}
	}

	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
