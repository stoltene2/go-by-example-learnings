package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	p := newPerson("me")
	fmt.Println("person", p)

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{age: 30, name: "reversed"})
	fmt.Println(&person{"hello", 1})
}
