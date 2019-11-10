package main

import "fmt"

/*
important: slice is not clear yet
*/

type person struct {
	name    string
	nidCard string
}

type vertex struct {
	x float64
	y float64
}

func main() {

	nums := []int{1, 2, 3}
	shoppingCartItemPrices := []float64{89.023, 7.03}

	persons := []person{
		person{name: "x", nidCard: "qw"},
		person{name: "x", nidCard: "qw"}}

	vertices := []vertex{vertex{x: 89.23, y: -023}, vertex{}}

	dynamicStructSlice := []struct {
		i int
		y bool
	}{
		{2, true},
		{20, false},
	}

	fmt.Println(nums)
	fmt.Println(shoppingCartItemPrices)
	fmt.Println(persons)
	fmt.Println(vertices)
	fmt.Println(dynamicStructSlice)

	fmt.Println()

	printSlice(dynamicStructSlice)

	// note the usage of '=' not ':=' what does it mean?
	dynamicStructSlice = dynamicStructSlice[:0]
	printSlice(dynamicStructSlice)

	// why are we getting 1 element here? seems like previous "assignment" hasn't changed the slice
	dynamicStructSlice = dynamicStructSlice[1:2]
	printSlice(dynamicStructSlice)

	dynamicStructSlice = dynamicStructSlice[1:]
	printSlice(dynamicStructSlice)

	var s []int // is not the same as 's := []int{}'
	fmt.Println(s == nil)
	fmt.Println(len(s), cap(s), s) // 0 0 []

	q := []int{}
	fmt.Println(len(q), cap(q), q)
}

func printSlice(v []struct {
	i int
	y bool
}) {
	fmt.Printf("len=%d cap=%d %v\n", len(v), cap(v), v)
}
