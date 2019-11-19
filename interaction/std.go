package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Enter: ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\t')

	if err == nil {
		fmt.Println("Entered input: ", input)
	} else {
		fmt.Println("Error occurred")
	}
}
