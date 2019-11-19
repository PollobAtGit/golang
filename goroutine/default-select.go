package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Printf("	tick!	")
		case <-boom:
			fmt.Printf("	BOOM!	")
			return
		default: // default select
			fmt.Printf("[....]")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
