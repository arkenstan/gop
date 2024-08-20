package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	arg     int
	message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("%d - %s", me.arg, me.message)
}

func fCE(arg int) (int, error) {
	if arg == 42 {
		return -1, &MyError{42, "Can't work with 42"}
	}
	return arg + 3, nil
}

func customError() {

	for _, i := range []int{5, 42} {

		if num, err := fCE(i); err != nil {

			var myE *MyError

			if errors.As(err, &myE) {
				fmt.Println("FOR ARG", myE.arg)
				fmt.Println("FOR MESSAGE", myE.message)
			} else {
				fmt.Println("Unknown error", err)
			}

		} else {

			fmt.Printf("Num %d\n\n", num)

		}

	}

}
