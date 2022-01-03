package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	length float64
}

type rectangle struct {
	length float64
	width  float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func (r rectangle) getArea() float64 {
	return r.length * r.width
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func getArea(s shape) float64 {
	return s.getArea()
}

func main() {
	mySquare := square{3}
	myRectangle := rectangle{3, 4}
	myTriangle := triangle{3, 4}

	fmt.Printf("Square. Length: %v. Area: %v\n", mySquare.length, getArea(mySquare))
	fmt.Printf("Rectangle. Length: %v. Width: %v. Area: %v\n", myRectangle.length, myRectangle.width, getArea(myRectangle))
	fmt.Printf("Triangle. Height: %v. Base: %v. Area: %v\n", myTriangle.height, myTriangle.base, getArea(myTriangle))
}
