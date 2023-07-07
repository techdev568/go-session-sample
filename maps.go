package main

import "fmt"

func Maps() {
	sample1()
	// we store data map -> key and value
	// no specific order
	// key -> value
	// var varName = map[type]type{values}

	var map1 = map[int]int{0: 1, 1: 2, 2: 3}
	fmt.Println(map1)
	fmt.Println("value", map1[1])
	//  varName := map[type]type{values}
	//var
	map2 := map[int]string{1: "apple", 2: "ball"}
	fmt.Println("map2", map2)

	var map3 = make(map[int]int)
	map3[123] = 100000
	fmt.Println("map2", map3)

	var map4 = make(map[int]string)
	map4[123] = "Alice"
	map4[12] = "Bob"
	map4[124] = "Don"
	map4[125] = "Elon"
	fmt.Println("map4", map4[12])
	map4[12] = "candy"
	fmt.Println("map4 new updated value", map4, map4[12])
	delete(map4, 123)
	fmt.Println("map4 new updated map", map4)
	// for loop -range
	for key, value := range map4 {
		fmt.Println("Loop", key, value)
	}

	// map refrence

	//mapVar1 <- map4

	map5 := map4
	fmt.Println("---", map5, map4)
	delete(map5, 12)
	fmt.Println("---", map5, map4)

}

// Funtions
// func funcName(input) return

func sample1() {

	a := sample2()
	fmt.Println("f1", a)
}

func sample2() int {
	fmt.Println("f2")
	a, b := sample3()
	fmt.Println("f2", a, b)
	return b
}

func sample3() (int, int) {
	fmt.Println("f3")
	return 1, 2
}
