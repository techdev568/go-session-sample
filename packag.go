package main

import (
	"fmt"

	maths "github.com/techdev568/golang_package"
)

// Pakage : fmt also one of the package comes with golang
func Pack() {
	fmt.Println("Hello")
	fmt.Println("Addition of two num", maths.Add(202, 203))
	fmt.Println("Addition of two num", maths.Sub(302, 203))
	fmt.Println("Addition of two num", maths.Mul(20, 20))
	//fmt.Println("Addition of two num", maths.Div(40, 20))
}

// go mod tidy -> to download all packages at a time
