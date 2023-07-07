package main

import (
	"fmt"
	"testing"
)

type addTest struct {
	arg1     int
	arg2     int
	expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{6, 1, 7},
	{3, 5, 8},
	{22, 33, 55},
}

func TestAdd1(t *testing.T) {

	fmt.Println("test add1")
	for _, val := range addTests {
		fmt.Println("val", val)
		if output := Add(val.arg1, val.arg2); output != val.expected {
			t.Error("error in add func")
		}
	}
}
