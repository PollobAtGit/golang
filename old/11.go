package main

import (
	"fmt"
)

func main() {

	message := make(chan string)

	// what's the point of '2'
	bufferredChannel := make(chan string, 2)

	go func() {
		message <- "x"

		bufferredChannel <- "x1"
		bufferredChannel <- "x2"
		bufferredChannel <- "x3"
	}()

	fmt.Println(<-message)

	fmt.Println(<-bufferredChannel)
	fmt.Println(<-bufferredChannel)
	fmt.Println(<-bufferredChannel)
}
