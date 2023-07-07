package main

import (
	"fmt"
	"runtime"
	"time"
)

func Selec() {
	fmt.Println("hello")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Two"
	}()
	fmt.Println("No of goroutines", runtime.NumGoroutine())
	for {
		select {
		case msg := <-c1:
			fmt.Println("hello select", msg)
		case msg1 := <-c2:
			fmt.Println("hello select", msg1)
		}
	}
}
