package main

import (
	"fmt"
	"math"
)

type square struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	return s.width * s.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 4,
	}
	s := square{
		height: 20,
		width:  10,
	}

	printArea(c)
	printArea(s)
	//fmt.Println(c.area())
	//fmt.Println(s.area())

}
