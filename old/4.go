package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")

	s := make([]string, 3)

	fmt.Println(s)

	fmt.Println(len(s))// 3

	s[0] = "index 1"
	s[1] = "index 2"
	s[2] = "index 3"

	fmt.Println(s)
	fmt.Println(make([]int, 5))
	fmt.Println(make([]bool, 3))

	fmt.Println(append(s, "t"))
	fmt.Println(s)

	qrt := make([]string, len(s))
	copy(qrt, make([]string, 3))
	copy(qrt, s)
	fmt.Println(qrt)

	t :=make([]int, 4)

	for i := 0 ; i < 4 ; i += 1 {
		t[i] = i + 10
	}

	fmt.Println()
	
	fmt.Println(t)
	fmt.Println(t[:1])
	fmt.Println(t[2:])
	fmt.Println(t[1:3])

	multiDimensionalSlice()
	workingWithMap()
}

func workingWithMap() {

	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20

	fmt.Println(m)

	v := m["a"]
	q := m["b"]

	fmt.Println(v)
	fmt.Println(q)

	fmt.Println(v + q)

	fmt.Println(len(m))

	delete(m, "a")

	fmt.Println(m)

	r := m["rr"]

	// r is zero indicating value not retrieved
	fmt.Println(r)

	w, y := m["rt"]

	fmt.Println(w, y)

	_, exists := m["y"]
	fmt.Println(exists)

	mm := map[string]int { "a" : 1, "b" : 2 }
	fmt.Println(mm)
	
	fmt.Println(map[int]float32 { 10 : 23.03 , 23 : 0.23 })
}

func multiDimensionalSlice() {
	
	m := make([][]int, 3)
	for i := 0 ; i < 3 ; i = i + 1 {
		n := i + 1
		m[i] = make([]int, n)
		for j := 0 ; j < n ; j = j + 1 {
			m[i][j] = i * j
		}
	}

	fmt.Println(m)
}

