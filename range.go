package main

import "fmt"

func rangeFn() {

	nums := []int{1, 2, 3}

	sum := 0

	for _, num := range nums {

		sum += num

		fmt.Println(num, sum)

	}

	kvs := map[string]string{"a": "apples", "b": "bananas"}

	for key, value := range kvs {
		fmt.Println(key, value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
