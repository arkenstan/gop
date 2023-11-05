package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2.0*r.width + 2.0*r.height
}

func (r circle) perimeter() float64 {
	return 2 * math.Pi * r.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func interfaces() {

	rect1 := rectangle{width: 20, height: 30}
	cir1 := circle{radius: 10}

	measure(rect1)
	measure(cir1)

}
