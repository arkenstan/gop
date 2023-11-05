package main

import "fmt"

func abVals() (int, int) {

	return 10, 20

}

func abcVals() (int, int, int) {
	return 10, 50, 30
}

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func funcs() {

	a, b := abVals()

	res := plus(a, b)
	fmt.Println(res)

	a, b, c := abcVals()

	res = plusPlus(a, b, c)
	fmt.Println(res)

}
