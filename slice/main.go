package main

import "fmt"

func main() {
	fmt.Println(make([]int, 5))
	fmt.Println([5]int{})

	fmt.Printf("Type %T\n", make([]int, 5)) //Type []int => slice
	fmt.Printf("Type %T\n", [...]int{})     // Type [0]int => array
	fmt.Printf("Type %T\n", []int{})        // Type []int => slice

	a := [...]int{1, 2, 3}

	fmt.Println(a[1:])
	fmt.Printf("Type %T\n", a)     // array
	fmt.Printf("Type %T\n", a[1:]) // slice

	checkSlicePointsToOriginalArray()

	fmt.Printf("a: %v, cap(a)= %d, len(a)= %d\n", a, cap(a), len(a))

	increasedSlice := increaseSliceSize(a[:])
	fmt.Printf("increasedSlice: %v, cap(increasedSlice)= %d, len(increasedSlice)= %d\n", increasedSlice, cap(increasedSlice), len(increasedSlice))
}

func checkSlicePointsToOriginalArray() {

	s := []string{}

	fmt.Printf("s: %v, cap(s)= %d, len(s)= %d\n", s, cap(s), len(s))

	s = append(s, "x")
	s = append(s, "y")
	s = append(s, "z")
	s = append(s, "a")

	fmt.Printf("s: %v, cap(s)= %d, len(s)= %d\n", s, cap(s), len(s))

	s = append(s, "b")

	fmt.Printf("s: %v, cap(s)= %d, len(s)= %d\n", s, cap(s), len(s))

	q := s[1:3]

	fmt.Println(q)

	q[0] = "A"

	fmt.Println(s) //[x A z a b]
}

func increaseSliceSize(originalSlice []int) []int {
	copiedSlice := make([]int, len(originalSlice), (cap(originalSlice)+1)*2)

	for i := range originalSlice {
		copiedSlice[i] = originalSlice[i]
	}

	return copiedSlice
}
