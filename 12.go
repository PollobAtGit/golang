package main

import (
	"time"
)

func square(doneSignaler chan bool) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
	}

	doneSignaler <- true
}

func main() {

	signaler := make(chan bool)

	go square(signaler)

	<-signaler
}
