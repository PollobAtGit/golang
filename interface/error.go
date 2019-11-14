package main

import (
	"fmt"
	"time"
)

type domainError struct {
	when time.Time
	what string
}

// 'Error() string' gets presidence over 'String() string'
func (de domainError) Error() string {
	return fmt.Sprintf("at %v, %s", de.when, de.what)
}

func (de domainError) String() string {
	return fmt.Sprintf("String(...): %v", de.what)
}

func main() {

	// at 2019-11-14 23:16:09.3429912 +0600 +06 m=+0.002000401, it didn't work!
	fmt.Println(domainError{when: time.Now(), what: "it didn't work!"})
	fmt.Println(getError())

	operateOnError()
}

func operateOnError() {

	e := getErrorIf(true)

	if e == nil {
		fmt.Println("no error occurred")
	} else {
		fmt.Println(e)
	}
}

func getError() domainError {
	return domainError{when: time.Now(), what: "runtime error"}
}

// what's the difference between 'error' & 'domainError'
func getErrorPointer() domainError {
	return domainError{when: time.Now(), what: "compilation error"}
}

func getErrorIf(is bool) *domainError {
	if is {
		return &domainError{time.Now(), "unknown error"}
	}
	return nil
}
