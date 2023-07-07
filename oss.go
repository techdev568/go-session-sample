package main

import (
	"bufio"
	"fmt"
	"os"
)

func OSS() {
	fmt.Println("hello")
	file, err := os.Open("samplefile.txt")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("File was opened sucuusfully")

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("There is a problem in scanner")
		}
	}

}
