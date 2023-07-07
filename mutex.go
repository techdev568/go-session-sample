package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu       sync.Mutex
	counters map[string]int // =>counters["abc"]=1
}

func (c *Counter) inc(name string) {
	fmt.Println("Counter name", name, c.counters[name])
	c.mu.Lock()
	c.counters[name]++
	defer c.mu.Unlock()
}

func Mutex() {
	fmt.Println("Hello")
	var wg sync.WaitGroup
	c := Counter{
		counters: map[string]int{"a": 0, "b": 0},
	}
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10)
	go doIncrement("a", 5)
	go doIncrement("b", 7)

	wg.Wait()
	fmt.Println("Counter", c.counters)
}
