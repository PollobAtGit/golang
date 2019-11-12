package main

import (
	"fmt"
	"math"
)

func main() {

	f := func(a, b float64) float64 { return math.Sqrt(a * b) }

	functionAsValue()
	fmt.Println(invoke(f))

	closure()

	fmt.Println(getInvokable()(89, 10))
}

func getInvokable() func(float64, float64) float64 {
	var s float64 = 56.2
	return func(a, b float64) float64 {
		return a + b + s
	}
}

func closure() {
	s := 0
	q := func(a int) {

		// 's' is in closure
		s += a
	}

	q(48)
	q(10)

	fmt.Println(s)
}

// note keyword 'func' after variable
func invoke(f func(float64, float64) float64) float64 {
	return f(0.23, 0.9)
}

func functionAsValue() {
	f := func(a, b float64) float64 {
		return math.Sqrt(a * b)
	}

	fmt.Println(f(8, 9.3))
}
