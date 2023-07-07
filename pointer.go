package main

import "fmt"

//a variable it holds the address
// we can store the address of the another varibale
func pointer() {
	// var VarName *Type
	//* -> give value varibale
	//& -> address of the variable
	var ptr *int // default value will be nil
	fmt.Println("hello", ptr)
	a := 10
	fmt.Println("hello", a)
	ptr = &a
	//fmt.Println("hello", ptr)
	fmt.Println("hello", &ptr)
	//fmt.Println("ptr value", *ptr)

	ptrf1(ptr)
	fmt.Println("ptrf", *ptr)
	fmt.Println("hello", a)

	//Pointer
}

func ptrf(_a int) {
	_a = 20
	fmt.Println("ptrf", _a)
}

func ptrf1(prt *int) {
	*prt = 20
	fmt.Println("ptrf", *prt)
}
