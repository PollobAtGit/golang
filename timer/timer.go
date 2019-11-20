package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const totalDuration time.Duration = 1e+10

func shouldExitMessageEmitter(st time.Time, sq chan bool) {
	if time.Now().Sub(st) >= totalDuration {
		sq <- true
	}
}

func tick() {

}

func readString(readerChannel chan string) {

	stdinReader := bufio.NewReader(os.Stdin)

	s, e := stdinReader.ReadString('\n')

	if e == nil {
		readerChannel <- s
	}
}

func checkIfShouldQuit(startTime time.Time, shouldQuit chan bool) {
	for range time.Tick(time.Second) {
		if time.Now().Sub(startTime) >= totalDuration {
			shouldQuit <- true
		}
	}
}

func takeQuiz() {

	shouldQuit := make(chan bool)
	readerChannel := make(chan string)

	startTime := time.Now()

	go checkIfShouldQuit(startTime, shouldQuit)

	go readString(readerChannel)

	isExiting := false

	for !isExiting {

		fmt.Printf("Enter input:\t")

		select {
		case msg := <-readerChannel:
			fmt.Println("You entered:\t", msg)
			go readString(readerChannel)
		case <-shouldQuit:
			fmt.Println("EXITTING!")
			isExiting = true
		}
	}
}

func main() {
	// tickExercise()
	takeQuiz()
}

func tickExercise() {

	t := time.Tick(time.Second)

	fmt.Println("Starting tick: \t", (<-t).Second())

	stdin := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("Enter input (current tick => %v):\t", (<-t).Second())

		s, e := stdin.ReadString('\n')

		if e == nil {
			fmt.Println("You entered:\t", s[:len(s)-1])
			// fmt.Printf("You entered:\t%v. Current Tick:\t%v", s[:len(s)-1], (<-t).Second())
		}

		if (<-t).Second() >= 30 {
			fmt.Println("TIMEOUT!")
			break
		}
	}
}
