package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type container2 struct {
	a   base
	b   base
	str string
}

func structEmbedding() {

	co2 := container2{a: base{num: 44}, b: base{num: 22}, str: "Another Name"}

	fmt.Printf("co2= {numA: %v, numB: %v, str: %v}\n", co2.a.num, co2.b.num, co2.str)

	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

}
