package main

import "fmt"

func conditions() {

	fmt.Println("========== CONDITIONS")

	fmt.Println("# IF ELSE")

	a := 10

	if a%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	fmt.Println("# SWITCH")

	var i int = 10

	switch i {
	case 1:
		fmt.Println("is 1")
	case 10:
		fmt.Println("is 10")

	}

}
