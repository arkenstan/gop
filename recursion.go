package main

import "fmt"

func factorial(num int) int {

	if num <= 0 {
		return 1
	}

	return num * factorial(num-1)

}

func recursion() {

	fmt.Println(factorial(5))

	var fibs func(n int) int

	fibs = func(n int) int {
		if n < 2 {
			return n
		}

		return fibs(n-1) + fibs(n-2)

	}

	fmt.Println(fibs(7))

}
