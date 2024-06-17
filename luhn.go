package main

import "fmt"

func luhnCC(cardNumbers []int) {

	parity := len(cardNumbers) % 2

	sum := 0

	for i := len(cardNumbers) - 2; i >= 0; i-- {
		if i%2 != parity {
			sum += cardNumbers[i]
		} else if cardNumbers[i] > 4 {
			sum += 2*cardNumbers[i] - 9
		} else {
			sum += 2 * cardNumbers[i]
		}
	}

	check := 10 - (sum % 10)

	fmt.Println("Is Valid", cardNumbers[len(cardNumbers)-1], check)

}
