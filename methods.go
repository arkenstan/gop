package main

import "fmt"

type rect2 struct {
	width, height int
}

func (r *rect2) area() int {
	return r.width * r.height
}

func (r rect2) perim() int {
	return 2*r.height + 2*r.height
}

func methods() {

	r := rect2{width: 40, height: 20}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perim())

	rp := &r

	fmt.Println("AREA: ", rp.area())
	fmt.Println("PERIMETER: ", rp.perim())

}
