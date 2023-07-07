package main

import "fmt"

func Calc() {
	fmt.Println("Hello")
	fmt.Println("calc", Add(3, 4))
	fmt.Println("calc", Sub(5, 4))

}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
