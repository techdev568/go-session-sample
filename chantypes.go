package main

import "fmt"

func ChanTypes() {
	fmt.Println("hello")

	//unbuffered channel
	msg1 := make(chan string)
	//buffered channel
	msg := make(chan string, 3)

	msg <- "msg"
	msg <- "channel"
	msg <- "channel1"
	fmt.Println("hello", <-msg)
	fmt.Println("hello", <-msg1)
	// fmt.Println("hello", <-msg)
}
