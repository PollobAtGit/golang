package main

import (
	"errors"
	"fmt"
	"math"
)

func invoke(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't process")
	}

	return arg + 1, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// why is &argError be returned for 'error'? but can't do same for 'argError'
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return -1, &argError{prob: "can't process negative numbers", arg: -1}
	}

	return math.Sqrt(n), nil
}

func main() {
	fmt.Println("sd")

	i, t := invoke(42)

	fmt.Println(i, t)

	ae := argError{arg: 10, prob: "r"}
	fmt.Println(ae.Error())

	fmt.Println()

	for _, i := range []int{42, 7} {
		fmt.Println(i)
		if n, e := invoke(i); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(n)
		}
	}

	for _, n := range []float64{-20, 20, -30} {
		if r, e := sqrt(n); e != nil {
			fmt.Println("error occurred: ", e)
		} else {
			fmt.Println("sqrt value: ", r)
		}
	}
}
