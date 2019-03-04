package main

import (
	"fmt"
	"math"
)

//define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

//area for circle
func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

//area for rect
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circ := Circle{x: 0, y: 0, radius: 5}
	rect := Rectangle{width: 10, height: 5}

	fmt.Println(getArea(circ))
	fmt.Println(getArea(rect))
}
