package main

import "fmt"

const (
	ab = iota
	_
	bd = iota + 4
	cd = iota + 2
	dc
)

type Size uint8

const (
	small  Size = 0
	medium Size = 1
	large  Size = 2
)

// Go lang , wont accept the dynamicchange of type of a variables

var ad int = 10

func IOTAS() {
	fmt.Println("Hello")
	fmt.Println("Hello", ab)
	fmt.Println("Hello", bd)
	fmt.Println("Hello", cd)
	fmt.Println("Hello", dc)

	fmt.Println("Hello enum", small)
	fmt.Println("Hello enum", medium)
	fmt.Println("Hello enum", large)
	fmt.Println("Hello enum", ad)

}
