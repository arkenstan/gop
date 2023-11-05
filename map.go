package main

import (
	"fmt"
	"maps"
)

func mapfn() {

	m := make(map[string]string)

	m["hello"] = "hello"

	m["some"] = "thing"

	fmt.Println(m)

	delete(m, "hello")

	fmt.Println(m)

	clear(m)

	fmt.Println(m)

	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("n1 === n2")
	}

}
