package main

import (
	"fmt"
	"os"
)

func main() {

	f, e := os.Create("new.txt")
	if e == nil {
		fmt.Println(f)
	} else {
		fmt.Println(e)
	}
}
