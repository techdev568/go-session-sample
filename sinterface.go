package main

import "fmt"

/**
What is interface?

Interface is a type
collection method signatures
it function name
it incluedes parameters
it includes retun values
***/

type animal interface {
	// method signatures
	walk()
	breath()
}

type lion struct {
	age int
}

type dog struct {
	age int
}

func (l *lion) walk() {
	l.age = 30
	//fmt.Println("lion walk", l)
	fmt.Println("lion walk", l)
}
func (l lion) breath() {
	fmt.Println("lion breath", l)
}

func (d dog) breath() {
	fmt.Println("dog breath", d)
}

func (d dog) walk() {
	d.age = 30
	//fmt.Println("lion walk", l)
	fmt.Println("dog walk", d)
}
func sinterface() {
	fmt.Println("hello")
	var a animal
	a = &lion{age: 22}
	a.walk()
	a.breath()

	d := dog{age: 5}
	callBreath(d)

}

func callBreath(a animal) {
	a.breath()
}
