package main

import (
	"fmt"
	"math"
)

type rect struct {
	height, width float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.height * r.width
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float32 {
	return 2 * r.height + 2 * r.width;
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	area() float32
	perim() float32
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	c := circle { radius : 89 }
	r := rect { height : 7, width : 8 }

	measure(c)

	fmt.Println()

	measure(r)
}
