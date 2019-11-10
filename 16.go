package main

import "fmt"

type person struct {
	name string // person name
	age  int    // person age
}

func main() {

	w := person{name: "s", age: 89}
	fmt.Println(w)
}
