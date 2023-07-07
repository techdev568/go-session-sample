package main

import "fmt"

//Varibale Declaration
//var x string = "hello"
//var varName varType = 10
var z = "sample"
var y int = 10
var temp, temp1 = 30, 40

//const constName type = value
const d int = 100

const e string = "hello func"

var f float32 = 2.0

var g = true

const fixed = false

var (
	a = "a"
	b = "b"
	c = "c"
)

func Varibales() {
	var x string = "hello"
	//varName := valueOf the variale
	sample1 := "Hello everyone"

	fmt.Println("Hello world", x, y, z, sample1, temp, temp1, d)
	fmt.Println("main func", e, f, g)
	f1(x)
}

func f1(_x string) {

	fmt.Println("f1 func", _x, d)
	f2()
}

//Letters
//Digits
//underscroes

// tempFile
// TempFile
// temp_file

// var -> golbally and locally anywhere in the file
// := we declare and use locally

//const -> declared and assigned with

func f2() {
	fmt.Println("--------------")
	fmt.Println("f2 func", e, f, g)
}
