package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2*math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("g: %T\n", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}


func main() {
	measure(rect{width: 3, height: 4})
	measure(circle{radius: 5})
}
