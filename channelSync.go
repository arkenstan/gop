package main

import "time"

func worker(done chan bool, done2 chan bool) {
	println("Working")
	time.Sleep(time.Second)
	<-done2
	println("Done")
	done <- true
}

func worker2(done2 chan bool) {
	println("Working 2")
	time.Sleep(2 * time.Second)
	println("Done 2")
	done2 <- true
}

func channelSyncTest() {
	done := make(chan bool)
	done2 := make(chan bool)

	go worker(done, done2)
	go worker2(done2)

	time.Sleep(3 * time.Second)

	<-done
}
