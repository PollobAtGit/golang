package main

import (
	"fmt"
	"image"
)

func main() {

	i := image.NewRGBA(image.Rect(0, 0, 150, 250))

	fmt.Println(i.Rect)
	fmt.Println(i.Bounds())
	fmt.Println(i.At(0, 0).RGBA())
	fmt.Println(i.At(0, 150).RGBA())

}
