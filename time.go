package main

import (
	"fmt"
	"time"
)

func Time() {
	fmt.Println("hello")

	t := time.Now()
	fmt.Printf("Now the time is %s\n", t)

	newTime := t.Add(time.Hour * 1)
	fmt.Printf("Now the time is %s\n", newTime)
	fmt.Printf("hello\n")

	newTime1 := t.Add(-time.Hour * 1)
	fmt.Printf("Now the time is %s", newTime1)

	newTime2 := t.Add(time.Minute * 12)
	fmt.Printf("Now the time is %s", newTime2)
}
