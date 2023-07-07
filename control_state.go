package main

import "fmt"

func ControlStatements() {
	fmt.Println("Hello world")
	//i := 1
	//sampleIF(i)
	//sampleSwitch(i)
	//sampleFor(i)
	sample()
}

var foo int = 10

func sample() {

	foo := 42 // <-- legal
	foo = 314 // <-- legal
	fmt.Println(foo, foo)
}
func sampleIF(i int) {
	// if
	if i == 1 {
		fmt.Println(" value One")
	} else if i == 2 {
		fmt.Println("value TWO")
	} else if i == 3 {
		fmt.Println("value Three")
	} else if i == 4 {
		fmt.Println("value FOUR")
	} else {
		fmt.Println("else block")
	}
	//
	fmt.Println("--------------------------")
	salary := 900
	pf := 300
	if salary >= 900 {
		if pf >= 200 {
			fmt.Println("PF suff")
		} else {
			fmt.Println("PF not suf")
		}
	} else {
		fmt.Println("No need PF")
	}

}

func sampleSwitch(i int) {
	fmt.Println("--------------------------")
	switch i {
	case 1, 5:
		fmt.Println("value 1")
	case 2, 4:
		fmt.Println("value is even")
	case 3:
		fmt.Println("value 3")
	default:
		fmt.Println("default")
	}
}

//>
//<
//<=
//>=
//!=
//==

//
//

func sampleFor(i int) {
	fmt.Println("--------------------------")
	for i < 10 {
		// state
		fmt.Println("value", i)
		i = i + 1
	}
	//
	for j := 1; j < 5; j++ {
		fmt.Println("value J", j)
	}
}
