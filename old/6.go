package main

import "fmt"

func main() {

	// multiple arguments
	fmt.Println(sum(10, 20, 30))

	a := []int { 89, 56 }
	
	// spread operator
	fmt.Println(sum(a...))

	f := intSequence()

	fmt.Println()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(factorial(5))
	fmt.Println(factorial(7))

	n := 23

	fmt.Println()

	fmt.Println("initial value: ", n)

	zerobyvalue(n)

	fmt.Println("after zerobyvalue: ", n)

	zerobypointer(&n)
	
	fmt.Println("after zerobypointer: ", n)

	// address
	fmt.Println("pointer: ", &n)

	fmt.Println(person { "bob", 12 })
	fmt.Println(person { "qrt", 89 })

	p := person { "q", 78 }

	fmt.Println(p)

	fmt.Println(createperson())
	fmt.Println(createpersonwithage(7))

	modifyperson(p)

	fmt.Println(p)

	// pass person pointer
	modifypersonwithpointer(&p)

	fmt.Println(p)
	
	// address?
	fmt.Println(&p)
}

func modifypersonwithpointer(p *person) {
	
	// get person pointer
	(*p).age = 8
	(*p).name = "88"
}

func modifyperson(p person) {
	p.name = "123"
	p.age = 9
}

func createpersonwithage(p_age int) person {
	p := person { age : p_age }
	p.name = "t"
	return p
}

func createperson() person {
	p := person { name : "ty" }
	p.age = 78

	return p
}

type person struct {
	name string
	age int
}

func zerobyvalue(n int) {
	n = 10
}

func zerobypointer(n *int) {
	*n = 20
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n - 1)
}

func intSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func sum(nums...int) map[int]int {
	s := map[int]int { }
	for i, num := range nums {
		_, exists := s[i]
		if !exists {
			if i == 0 {
				s[i] = num
			} else {
				s[i] = num + s[i - 1]
			}
		}		
	}
	
	return s
}
