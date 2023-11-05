package main

import "fmt"

func arrays() {

	fmt.Println("========== ARRAYS")

	// count := 5
	var a [5]int

	a[2] = 5

	fmt.Println(a)

	var twoD [3][4]int

	for i := 0; i < 3; i++ {

		for j := 0; j < 4; j++ {

			twoD[i][j] = i + j
		}

	}

	fmt.Println(twoD)

}
