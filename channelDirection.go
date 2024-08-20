package main

func ping(tx chan<- string, msg string) {
	tx <- msg
}

func pong(rx <-chan string, forward chan<- string) {
	msg := <-rx
	forward <- msg
}

func channelDirectionTest() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Message passed")
	pong(pings, pongs)

	println("MESSAGE", <-pongs)

}
