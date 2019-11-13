package main

import "fmt"

// 'type' don't have to be complex type all the time
type bigInt int64

type vertex struct {
	latitude, longitude float64
}

func (b *bigInt) add(cv bigInt){

	// dereferencing pointer to update value
	(*b) += cv
}

// try returning bigInt
func (b bigInt) abs() int64 {
	if (b < 0) {
		return int64(-b)
	}

	return int64(b)
}

func (v *vertex) add(latitude, longitude float64) {
	v.latitude += latitude
	v.longitude += longitude
}

func (v *vertex) remove(latitude, longitude float64) {
	v.latitude -= latitude
	v.longitude -= longitude
}

func (v vertex) multiplier() float64 {
	return v.latitude * v.longitude
}

func (v vertex) divider() float64 {
	return v.latitude / v.longitude
}

type car struct {
	model, manufacturer string
	price float64
}

// this is a method. so it has value argument which accepts both pointer and value
func (c *car) increasePrice(p float64) {
	c.price += p
}

// will except only pointers to 'car'
func toString(c *car) string {

	// string default value is not nil rather empty string
	// *string can be nil though

	if c.model == "" {
		return "model not defined yet"
	}

	return c.model
}

func main() {

	v := vertex{}
	v.add(0.23, -96.23)

	fmt.Println(v)              // {0.23 -96.23}
	fmt.Println(v.multiplier()) //-22.132900000000003
	fmt.Println(v.divider())

	var d bigInt

	d = 56

	d.add(8)
	d.add(10)

	fmt.Println(d)

	var bi bigInt

	bi = -9

	fmt.Println(bi.abs())

	c := car{}
	cr := &car{}

	// increasePrice accepts both: pointer and value
	c.increasePrice(48)
	cr.increasePrice(48)

	fmt.Println(c)
	fmt.Println(*cr)

	fmt.Println(toString(cr))
}
