package main

import "fmt"

func main() {
	mapZeroValue()

	// following fails because 'm' is nil
	// m["x"] = vertex{}

	initializeMap()
	addAndGet()
	literalInitialization()
	deleteElementByChecking()
}

func deleteElementByChecking() {
	q := map[string]vertex{
		"a": {latitude: 78, longitude: 9},
		"b": {latitude: 99, longitude: -0.23},
	}

	fmt.Println(q)

	key := "a"

	_, ok := q[key]

	if ok == true {
		delete(q, key)
	}

	fmt.Println(q)
}

func literalInitialization() {
	q := map[string]vertex{}

	fmt.Println(q)
	fmt.Println(q == nil) // false

	p := map[string]vertex{
		"t": vertex{latitude: 1, longitude: 2},
		"y": vertex{latitude: 3, longitude: 4},
	}

	fmt.Println(p)

	// defining elements based on shape without mentioning type explicitly
	r := map[string]vertex{
		"a": {latitude: 78, longitude: 9},
		"b": {latitude: 7, longitude: 90},
		"c": {latitude: -0.23, longitude: -.59},
	}

	fmt.Println(r)
}

func addAndGet() {
	m["x"] = vertex{}
	m["y"] = vertex{latitude: 56, longitude: 78}

	fmt.Println(m) // map[x:{0 0} y:{56 78}]
	fmt.Println(m["x"])

	v, e := m["k"]

	// map returns default (zero value) of that type if not found
	// in this case need to depend on 'error'

	fmt.Println(v, e) //{0 0} false
}

type vertex struct {
	latitude, longitude float64
}

// creates a map with zero value (nil)
var m map[string]vertex

func mapZeroValue() {
	fmt.Println(m)        // map[]
	fmt.Println(m == nil) // true
}

func initializeMap() {
	m = make(map[string]vertex)
	mapZeroValue() // map[] ; m == nil returns false
}
