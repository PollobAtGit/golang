package main

import "fmt"

type vertex struct {
	X int
	Y int
}

var (
	v   = vertex{X: 8, Y: 99}
	vv  = vertex{}
	vvv = vertex{X: 7}
)

func main() {
	fmt.Println(vertex{1, 8})

	v := vertex{89, 7}
	vp := &v

	vp.X = 1
	vp.Y = 9

	fmt.Println(v)

	updateVertexProperties(vp)
	fmt.Println(v)

	updateVertexPropertiesByValue(v)
	fmt.Println(v)

	fmt.Println(v, vv, vvv)

	primes := []int{2, 3, 5, 7}

	fmt.Println(primes)

	// first prime
	fmt.Println(primes[:1])

	// last prime
	fmt.Println(primes[3:])

	// all primes starting from 2nd
	fmt.Println(primes[2:])
}

// similar to other languages the struct is passed as value
func updateVertexPropertiesByValue(v vertex) {
	v.X = 99
	v.Y = 100
}

func updateVertexProperties(v *vertex) {
	v.X = 70
	v.Y = 90
}
