package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c <- "x"
	}()

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(time.Second):
		fmt.Println("time-out")
	}
}
