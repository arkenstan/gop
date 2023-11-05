package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func pointers() {
	i := 1

	fmt.Println("Initial:", i)

	zeroVal(i)
	fmt.Println("Zero Val:", i)

	zeroPtr(&i)
	fmt.Println("Zero Ptr:", i)

	fmt.Println("pointer:", &i)

}
