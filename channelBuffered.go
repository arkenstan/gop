package main

import (
	"fmt"
	"time"
)

func channelBufferedTest() {

	message := make(chan string, 2)

	go func() {
		fmt.Println("Pehle")
		fmt.Println("Badme")
		for i := range 3 {
			message <- fmt.Sprintf("Ping %d", i)
		}

	}()

	fmt.Println("Message", <-message)
	fmt.Println("Message", <-message)
	// fmt.Println("Message", <-message)

	time.Sleep(time.Second)

	msg := <-message

	// msg := "Hello"

	fmt.Println("msg", msg)

}
