package main

import "fmt"

func main() {

	fmt.Println("hello world!")

	rangeFunctionality()

	fmt.Println(plus(10, 20))
	fmt.Println(plusPlus(20, 30, 40))
	fmt.Println(returningTuple(10))
}

func returningTuple(a int) (int, int) {
	return a + 10, a + 20
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func rangeFunctionality() {
	nums := [...]int { 1, 1, 2, 3, 5, 8, 13 }
	s := 0
	for _, num := range nums {
		s += num
	}

	fmt.Println(s)

	indexSum := 0
	for i, n := range nums {
		indexSum += i + n
	}

	fmt.Println(indexSum)

	m := map[int]int { }
	for i, _ := range [...]int { 23, 5 } {
		_, exists := m[i]
		if !exists {
			m[i + 20] = i + 10
		}
	}

	fmt.Println(m)

	fmt.Println()
	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println()
	for i, v := range "golang" {
		fmt.Println(i, v)
	}
}