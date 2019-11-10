package main

import (
	"fmt"
	"time"
)

func goRoutine(identifier string) {
	for i := 0; i < 3; i++ {
		fmt.Println("From: ", identifier, " => iteration: ", i)
	}
}

func main() {
	goRoutine("x")

	go goRoutine("y")
	go func(from string) {
		fmt.Println(from)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("done with main thread")
}
