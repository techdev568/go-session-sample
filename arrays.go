package main

import "fmt"

func arrays() {

	//sampleArray()
	sampleSlice()
}
func sampleSlice() {
	//Type1
	var slic = []string{"a", "b", "c"}
	fmt.Println("Array values", slic)
	//Type2 make
	var slice1 = make([]int, 5, 5) // capcity will not be < length
	fmt.Println("Array values", slice1)
	fmt.Println("Array values", len(slice1))
	fmt.Println("Array values", cap(slice1))
	// Possibility of inserting the value into slice
	var arr1 = [5]int{1, 2, 3, 4, 5}
	// arr1[start:end]
	slice3 := arr1[2:3] // length = end-start , // capacity = length of the array -start of slice
	fmt.Println("Array values", slice3, len(slice3), cap(slice3))
	slice5 := []int{11, 22, 33, 44}
	slice6 := append(slice3, slice5...)
	slice5 = append(slice5, 22, 33)
	fmt.Println("append values", slice6, len(slice6), cap(slice6))
	fmt.Println("Array values", slice5, len(slice5), cap(slice5))
}
func sampleArray() {
	// var keyword varName = []type{values}
	var arr = []int{1, 2}
	// var keyword varName = [length]type{values}
	var arr2 = []int{}            // Not Initialized
	var arr1 = [4]int{1, 2}       // partial Initialization
	var arr4 = [4]int{4, 3, 2, 1} //Fully init
	// index 0 and 1 remaning 0's
	arr6 := [4]string{"alice", "bob", "candy", "dan"}

	var arr5 = [4]string{"4", "3", "2", "1"} //Fully init

	fmt.Println("Array values", arr)
	fmt.Println("Array values", arr1)
	fmt.Println("Array values", arr2)
	fmt.Println("Array values", arr5, len(arr5))
	fmt.Println("Array values", arr4)

	//Arrays lenght of 4
	arr1[2] = 20
	arr1[3] = 30

	//Arrays lenght of 4
	arr1[2] = arr4[0]
	arr1[3] = 30

	fmt.Println("Array after update", arr1)
	fmt.Println("Array after update", arr1)

	//for -> range
	// for index,val range array
	for i, val := range arr6 {
		fmt.Println("index", i)
		fmt.Println("value", val)
	}

	for i := 0; i < len(arr6); i++ {
		fmt.Println("2nd loop", arr6[i])
	}
}
