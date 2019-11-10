package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	var i64 int64 = 10
	var i32 int32 = int32(i64)
	var i16 int16 = int16(i32)
	var i8 int8 = int8(i16)

	var ui64 uint64 = 30
	var ui32 uint32 = uint32(ui64)
	var ui16 uint16 = uint16(ui32)
	var ui8 uint8 = uint8(ui16)

	var f64 float64 = -23.023
	var f32 float32 = float32(f64)

	fmt.Println(i64, i32, i16, i8)
	fmt.Println(ui64, ui32, ui16, ui8)

	fmt.Println(byte(23))

	// rune us a shorthand for int32
	fmt.Println(rune(-23))

	fmt.Println(f64, f32)

	i := 2
	q := 0.23
	b := false

	fmt.Println()

	fmt.Printf("[i] is of type %T\n", i)
	fmt.Printf("[q] is of type %T\n", q)
	fmt.Printf("[b] is of type %T\n", b)

	const qt = 89
	const qtStr = "rt"

	fmt.Printf("[qt] is of type %T\n", qt)
	fmt.Printf("[qtStr] is of type %T\n", qtStr)

	fmt.Println(sum(10))
	fmt.Println(sum(20))

	sumTillThousand()

	fmt.Println()

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(4, 2, 10))
	fmt.Println(pow(2, 3, 100))

	fmt.Println()
	showOsVersion()

	fmt.Println()
	saturdayInDays()

	fmt.Println()

	fmt.Println(eligibilityDetector(10))
	fmt.Println(eligibilityDetector(20))
	fmt.Println(eligibilityDetector(-20))

	fmt.Println(testingDefer(89.23))
	fmt.Println(testingDefer(-0.3))

	fmt.Println()

	fmt.Println("returned value ", stackingDefer(3, 56))

	fmt.Println()

	fmt.Println("returned value ", stackingDefer(5, 10))

	pointerProcessor()
}

func sum(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += i
	}

	return s
}

func sumTillThousand() {
	q := []int{}
	s := 1
	for s < 100 {
		q = append(q, s)
		s += s
	}

	fmt.Println(q)
}

func pow(x, n, lim float64) float64 {
	if q := math.Pow(x, n); q < lim {
		return q
	}
	return lim
}

func showOsVersion() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("current os is ", runtime.GOOS)
	case "darwin":
		fmt.Println("darwin operating system")
	}
}

func saturdayInDays() {
	today := time.Now().Weekday()

	switch time.Monday {
	case today:
		fmt.Println("today ...")
	case today + 1:
		fmt.Println("tomorrow ...")
	case today + 2:
		fmt.Println("in the future ...")
	default:
		fmt.Println("don't know ...")
	}
}

func eligibilityDetector(n int) string {

	// same as 'switch true ...'
	// idiomatic usage instead of long if-else
	switch {
	case n < 0:
		return "can't process negative"
	case n > 0 && n < 10:
		return "not sure if number can be processed"
	case n > 10 && n < 20:
		return "can be processed but might cause issue"
	case n > 50:
		return "can be processed"
	default:
		return "undetermined"
	}
}

func testingDefer(n float64) float64 {

	// will be executed after enclosing function completes execution
	defer fmt.Println("stopped processing testingDefer")

	return math.Sqrt(n)
}

func stackingDefer(n int, m float64) float64 {

	for i := 0; i < n; i++ {
		defer fmt.Println("executing deferred statement ", i)
	}

	fmt.Println("Returning from stackingDefer")

	return math.Sqrt(m)
}

func pointerProcessor() {

	a, b := 10, 30

	// gets address to those variables
	aPtr, bPtr := &a, &b

	fmt.Println(aPtr, bPtr)
	fmt.Println(*aPtr, *bPtr)

	*bPtr, *aPtr = 89, 100

	fmt.Println(*aPtr, *bPtr)
}
