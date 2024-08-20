package main

import (
	"errors"
	"fmt"
)

func f(i int) (int, error) {
	if i == 43 {
		return -1, errors.New("Cannot work with 43")
	}
	return i + 3, nil
}

var errorOutOfTea = fmt.Errorf("No more tea available")
var errorPower = fmt.Errorf("cannot boil water")

func makeTea(num int) error {
	if num == 2 {
		return errorOutOfTea
	} else if num == 4 {
		return fmt.Errorf("Making tea: %w", errorPower)
	}
	return nil
}

func errorLearn() {

	for _, i := range []int{1, 45, 43} {

		if num, err := f(i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("FOR %d: %d \n", i, num)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, errorOutOfTea) {
				fmt.Println("we should buy some more tea.")
			} else if errors.Is(err, errorPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("Unknown error %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!!")
	}

}
