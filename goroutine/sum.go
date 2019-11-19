package main

import (
	"fmt"
	"time"
)

func sum(l []int, c chan<- int, i int) {

	s := 0
	for _, n := range l {
		time.Sleep(time.Second)
		// fmt.Printf("For slice [%v - %v]\n", l[0], l[len(l)-1])
		fmt.Printf("%v ", i)
		s += n
	}

	c <- s
}

func parallelSummation() {

	l := make([]int, 50)

	for i := 0; i < 50; i++ {
		l[i] = i + 100
	}

	s := make(chan int)

	go sum(l[:len(l)/3], s, 1)
	go sum(l[len(l)/3:len(l)/2], s, 1)
	go sum(l[len(l)/2:], s, 3)

	o, w, q := <-s, <-s, <-s

	fmt.Println(o, w, q)
	fmt.Println(o + w + q)
}

func send(c chan int) {

	for i := range [20]int{} {
		time.Sleep(time.Second)
		fmt.Println("done with processing")
		c <- i
	}

	close(c)
}

func sendingDataContinuously() {

	c := make(chan int)

	go send(c)

	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	// parallelSummation()
	sendingDataContinuously()
}
