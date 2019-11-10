package main

import "fmt"

// why can't read channel here
func messageSender(messageChannel chan<- string) {
	messageChannel <- "x"
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {

	messageChannel := make(chan string, 1)

	messageSender(messageChannel)

	fmt.Println(<-messageChannel)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "ft")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
