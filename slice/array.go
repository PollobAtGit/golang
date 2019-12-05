package main

import "fmt"

// following doesn't  update value in original array
func updateArray(a *[2]int) {
	a[0] = 20
}

// a mew copy of the original array is passed here
func modifyArray(a [2]int) {

	// change in a doesn't change original array as value has been passed by value
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 10
	}

	fmt.Println(a)
}

func zeroValueOfSlice() {

	// here the slice is initialized but has no elements	
	fmt.Println([]int{} == nil) // false

	// declares an uninitialized variable	
	var slice []int

	fmt.Println(slice == nil) // true

	// from following comparisons it's difficult to figure if the slice was nil
	fmt.Println(len([]int{}), len(slice))// 0 0
	fmt.Println(cap([]int{}), cap(slice))// 0 0

	// following is not a valid operation as arrays are not nil-lable
	// fmt.Println([2]int{} == nil)	
}

func sliceComparison() {

	// invalid! slice can only be compared to nil
	// fmt.Println([]int{ 1 } == []int { 1 })

	// b := []int{ 1 }

	// fmt.Println(b == b[:])
}

func sliceFromArray() {
	
	a := [2]int{}
	slice := a[:]

	// comparison is valid  proves it's an array
	fmt.Println(slice == nil)
	fmt.Printf("Type of slice is %T\n", slice)

	// invalid ! as arrays are not nil-lable. cannot convert nil to type [2]int
	// fmt.Println(a == nil)
}

func slicingWholeArray() {
	
	a := [5]string{ "one", "two", "three", "four", "five" }
	
	slice := a[:]

	fmt.Println(slice)

	twoToFour := slice[1:4]
	
	fmt.Println(twoToFour)

	// changes twoToFour, slice and a
	twoToFour[1] = "hundred"

	fmt.Println(twoToFour, slice, a)
}


func sliceToIncreaseCap() {
	
	a := [3]int { 10, 29, 33 }

	slice := a[:]

	fmt.Println(cap(a), cap(slice))
	
	slice = slice[:cap(slice)]

	fmt.Println(cap(a), cap(slice))

	// will cause panic
	// slice = slice[:10]
}

func main() {
	
	a := [2]int{}

	for _, n := range a {
		fmt.Println(n == 0)	
	}


	fmt.Println(a)

	modifyArray(a)

	fmt.Println(a)

	updateArray(&a)
	
	zeroValueOfSlice()

	sliceComparison()

	sliceFromArray()

	slicingWholeArray()

	sliceToIncreaseCap()
}

