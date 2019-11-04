package main

import "fmt"

type car struct {
	model string
	price float32
}

func (c *car) increase_price_by(increase_amount float32) {
	c.price += increase_amount
}

func (c *car) update_name_by(n string) {
	c.model = n
}

// passed by value: so name won't be updated
func (c car) update_name(n string) {
	c.model = n
}

type rect struct {
	height, width int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2 * r.height + 2 * r.width
}

func main() {

	c := car { model : "xdrt", price : 23.023 }
	
	c.increase_price_by(89)
	c.update_name("bty")
	
	fmt.Println(c)

	c.update_name_by("56")
	fmt.Println(c)

	r := rect { height : 56, width : 78 }
	fmt.Println(r.area())
	fmt.Println(r.perim())
}