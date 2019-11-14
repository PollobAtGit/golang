package main

import "fmt"

type ipAddr [4]byte

func (i ipAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {

	m := map[string]ipAddr{
		"localhost":  {127, 0, 0, 1},
		"google-dns": {8, 8, 8, 8},
	}

	// fmt.Println(m)
	fmt.Println(m == nil)

	for k, v := range m {
		fmt.Printf("%v : %v\n", k, v)
	}
}
