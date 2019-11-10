package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		a <- "x"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		b <- "y"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-a:
			fmt.Println("received: ", msg)
		case msgO := <-b:
			fmt.Println("received: ", msgO)
		}
	}
}
