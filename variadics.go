package main

import "fmt"

func sumFn(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func variadics() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum := sumFn(nums...)

	fmt.Println(sum)

}
