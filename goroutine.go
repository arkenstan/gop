package main

import (
	"fmt"
)

func fStr(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func goRoutineTest() {

	fStr("Direct")

	go fStr("Go Routine")

	go func(msg string) {
		fmt.Println(msg)
	}("InLINE Going")
	// time.Sleep(time.Second)
	//

	//
	fmt.Println("done")
}
