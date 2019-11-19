package main

import (
	"fmt"
	"time"
)

func sendingRequest(c chan int, q chan bool) {

	for {
		select {
		case msg := <-c:
			fmt.Printf("received packet: %v\n", msg)
		case <-q:
			fmt.Println("quiting")
			return
		}
	}
}

func dataSender(c chan int, q chan bool) {

	for i := range [10]byte{} {
		time.Sleep(time.Second)
		c <- i
	}

	q <- true
}

func main() {

	c := make(chan int)
	q := make(chan bool)

	go dataSender(c, q)
	sendingRequest(c, q)

	fmt.Println("done!")
}
