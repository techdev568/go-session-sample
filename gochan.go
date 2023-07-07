package main

import (
	"fmt"
	"time"
)

// Multiple goroutines ->
func GoChan() {
	// pion-webrtc -> not able see video -> after
	//var a chan int // type1
	a := make(chan int)

	//fmt.Println("hello", a)
	//fmt.Println("hello", b)
	go receive(a)
	fmt.Println("send")
	go send(a)
	// fmt.Println("receive", val)
	time.Sleep(3000)
}

func send(a chan int) {
	//time.Sleep(2000)
	a <- 10
	fmt.Println("send")
	// val := <-a
	// fmt.Println("receive", val)
}

func receive(a chan int) {
	val := <-a
	fmt.Println("receive", val)
}
