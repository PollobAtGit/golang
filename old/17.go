package main

import (
	"fmt"
	"math"
	"math/rand"
)

func divide(p, q float64) (r float64) { // named return. note the variable declaration

	// this variable will be bound with return argument because of 'naked return'
	r = p / q

	// 'naked return'
	return
}

// interesting! good
func add(x, y int, p, q float64) (int, float64) {
	return x + y, p + q
}

// unused package level variable but compiler is not complaining
// var for multiple variable declaration
var cSharp, java, golang string

// shorthand to initialize variable doesn't work outside function scope
// we := 65

func main() {
	fmt.Println("my favorite number is: ", rand.Intn(10))
	fmt.Println("sqrt(n): ", math.Sqrt(12))

	i, f := add(10, 20, 0.2, 56)

	fmt.Println(i, f)
	fmt.Println(divide(89.3, 7))

	// var can be used for multiple variable declaration
	var mn, tu string
	fmt.Println(mn, tu)

	// multiple variable declaration with variable initialization
	var p, q, r, s, t int = 10, 20, 30, 40, 50
	fmt.Println(p, q, r, s, t)
}
