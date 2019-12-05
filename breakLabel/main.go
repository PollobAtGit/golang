package main

import (
	"fmt"
)

func workWithBreakLabel(fruits []string) {
ExitImmediately:
	for _, fruit := range fruits {
		if fruit == "" {
			fmt.Println("Exiting")
			break ExitImmediately
		} else {
			fmt.Println(fruit)
		}
	}
}

func main() {
	workWithBreakLabel([]string{"apple", "orange", "", "coconut"})
	fmt.Println()
	workWithBreakLabel([]string{"apple", "orange", "coconut"})
}
