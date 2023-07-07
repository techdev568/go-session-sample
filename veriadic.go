package main

import "fmt"

// veriadic : we give the any number of arguments to function
// ex fmt.Println()
// It will executer lastly in function
func Veriadic() {
	fmt.Println("Hello")
	arr4 := []int{4, 3, 2, 1}
	sum(arr4...)
	closure()
	fact := factorial(6)
	fmt.Println("fact", fact)
	panicS()
	sf1()
}
func panicS() {
	fmt.Println("panicS")
	//panic("error")
}

func sf1() {
	defer func() {
		fmt.Println("sf1 defer")
		c := 2 + 3
		fmt.Println("c", c)
	}()
	defer fmt.Println("another defer")
	defer fmt.Println("d2 defer")

	fmt.Println("sf1")
}

// veriadic function
func sum(args ...int) {
	fmt.Println("values:: ", args)
	total := 0
	for i, v := range args {
		fmt.Println("values:: ", i, v)
		total += v // total = total + v
	}
	fmt.Println("values:: ", total)

}

//Closure

func closure() {
	add := func(x, y int) int {
		fmt.Println("closure")
		return x + y
	}
	fmt.Println("closure", add(1, 2))
}

// 4*3*2*1
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
