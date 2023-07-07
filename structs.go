package main

import "fmt"

type Person struct {
	//Any niumber of data with different datatypes
	name string
	age  int
}

func StructSample() {
	// struct collection of data with different datatypes

	//declration

	//type structName struct {}

	p := Person{name: "alice"}
	p.age = 35
	fmt.Println("struct", p)
	strf(p)
	fmt.Println("struct after update", p)
	p.ff2()
	fmt.Println("struct method update", p)

}

func strf(p Person) {
	p.age = 45
	fmt.Println("strf IN", p)
}

//function
func ff1() {
	fmt.Println("ff1")
}

// Method
func (p *Person) ff2() {
	fmt.Println("ff1", p)
	p.name = "Bob"
	fmt.Println("ff1 afer name edit", p)
}
