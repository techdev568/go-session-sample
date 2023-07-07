package main

import "fmt"

func Goto() {
	fmt.Println("Hello")

	goto FINISH
	fmt.Println("Hello")

FINISH:
	fmt.Println("Hello action")
}
