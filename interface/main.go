package main

import (
	"fmt"
	"math"
)

type abs interface {
	abs() float64
}

type vertex struct {
	latitude, longitude float64
}

func (v *vertex) abs() float64 {
	return math.Abs(v.latitude + v.longitude)
}

type rectangle struct {
	height, width float64
}

func (r rectangle) abs() float64 {
	return math.Abs(r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) String() string {
	return fmt.Sprintf("Circle (%v radius)", c.radius)
}

func (c circle) abs() float64 {
	return math.Abs(c.radius)
}

type home struct {
	owner, address string
	price          float64
}

func (h home) String() string {
	return fmt.Sprintf("Home (Owner %v)", h.owner)
}

func main() {
	interfacePrimer()
	emptyInterface()
	typeAssertion()
	usingTypeSwitches()

	fmt.Println(circle{})
	fmt.Println(home{owner: "tag"})
}

func usingTypeSwitches() {
	fmt.Println(getFoundTypeToString(circle{}))
	fmt.Println(getFoundTypeToString(rectangle{}))
}

func getFoundTypeToString(a abs) string {
	// gets type
	switch a.(type) {

	// circle is a type name
	case circle:
		return "circle"
	case rectangle:
		return "rectangle"
	default:
		return "not found"
	}
}

func typeAssertion() {
	var s interface{} = "hello"

	v, sOk := s.(string)
	w, wOk := s.(int64)

	fmt.Println(v, sOk) // hello true
	fmt.Println(w, wOk) // 0 false

	fmt.Println(assertRectangle(rectangle{}))
	fmt.Println(assertCircle(circle{}))

	fmt.Println()

	fmt.Println(assertCircle(rectangle{}))
	fmt.Println(assertRectangle(circle{}))
}

func assertCircle(a abs) bool {
	_, ok := a.(circle)
	return ok
}

func assertRectangle(a abs) bool {
	_, ok := a.(rectangle)
	return ok
}

// how can empty interface be useful
func emptyInterface() {
	var i interface{}
	fmt.Println(i) // <nil>
}

func interfacePrimer() {
	var a abs

	fmt.Println(a == nil) // nil

	// following will through runtime exception because there's nothing associated with the variable
	// fmt.Println(a.abs())

	v := vertex{latitude: 9.23, longitude: -98}

	fmt.Println(v)

	// interface can take only pointer?
	a = &v

	fmt.Println(a.abs())

	r := rectangle{height: -9.023, width: -0.123}

	fmt.Println(r)

	// for 'rectangle' abs() is defined for value type. so we can assign value
	// directly rather than pointer

	a = r

	fmt.Println(a.abs())
}
