package main

import "fmt"

type shape interface {
	GetArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	sq := square{sideLength: 2}
	printArea(sq)

	tri := triangle{base: 2, height: 3}
	printArea(tri)
}

func printArea(s shape) {
	fmt.Println("Area is", s.GetArea())
}

func (s square) GetArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) GetArea() float64 {
	return 0.5 * t.base * t.height
}
