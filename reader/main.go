package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	s := strings.NewReader("hello! reader")
	b := make([]byte, 8)

	for {
		n, e := s.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, e, b)
		fmt.Printf("b[:n] = %q\n\n", b[:n]) // %q may convert slice to it's string representation

		if e == io.EOF {
			break
		}
	}
}
