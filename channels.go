package main

import (
	"fmt"
)

func channelTest() {

	message := make(chan string)

	go func() {
		fmt.Println("Pehle")
		fmt.Println("Badme")
		message <- "Ping"
	}()

	msg := <-message

	// msg := "Hello"

	// time.Sleep(time.Second)

	fmt.Println("msg", msg)

}
