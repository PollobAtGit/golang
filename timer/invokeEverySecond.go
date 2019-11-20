package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("Hello")
		}
	}()

	time.Sleep(5000)
	fmt.Println("Executed")
}
