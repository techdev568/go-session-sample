package main

import (
	"errors"
	"fmt"
	"os"
)

func Errorr() {

	_, err := os.Open("samplefile1.txt")
	if err != nil {
		smplerr := errors.New("There is problem in opening a file becuase not exist")

		fmt.Println("Error", smplerr)
	} else {

	}

	sum, err1 := add(-1, 5)

	if err1 != nil {
		fmt.Println("Error", err1)
	} else {
		fmt.Println("Sum", sum)
	}

}

func add(a int, b int) (int, error) {
	// T || T  = T
	// T || F  = T
	// F || T  = T
	// F || F  = F

	// T && T  = T
	// T && F  = F
	// F && T  = F
	// F && F  = F
	// !=
	// ==
	if a > 0 && b > 0 {
		return a + b, nil
	} else {
		err := errors.New("We will not accept negative values")
		return a, err
	}
}
