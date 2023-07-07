package main

import (
	"fmt"
	"sync"
)

func WG() {
	fmt.Println("Hello")
	var wg sync.WaitGroup
	//a := make(chan int)
	wg.Add(2)
	go start(&wg, 4)
	go start(&wg, 6)
	wg.Wait()
}

func start(wg *sync.WaitGroup, j int) {
	fmt.Println("start")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d*%d=%d\n", j, i, j*i) // 1*1=1
	}
	defer wg.Done()
}
