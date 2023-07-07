package main

import (
	"fmt"
	"sort"
	"sync"
)

//var ch chan int

func main() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go f22(ch)
	// go f11(ch, &wg) // goroutine
	// wg.Wait()
	// fmap()
	polindrome(1221)
}

func polindrome(num int) {
	fmt.Println("pol func", num)
	copynum := num
	var res int

	for num > 0 {
		i := num % 10

		res = res*10 + i
		num = num / 10
	}
	fmt.Println(copynum, res)
	if copynum == res {
		fmt.Println("polindrome")
	} else {
		fmt.Println("Not polindrome")
	}
}
func f11(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("f11")
	for i := 1; i < 6; i++ {
		ch <- i
	}
	wg.Done()

}
func f22(ch chan int) {
	for {
		select {
		case val := <-ch:
			if val == 5 {
				close(ch)
			}
			fmt.Println("hello f22", val)
		}

	}
	// Print all values at a time
	// for  v := range ch {
	// 	fmt.Println("hello f22", v)
	// }
}

func fmap() {
	m := make(map[string]int)
	m["C"] = 100
	m["A"] = 101
	m["K"] = 107
	m["D"] = 102
	fmt.Println("hello fmap", m)
	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println("before sort", keys)

	sort.Strings(keys)
	fmt.Println("after sort", keys)
	for _, k := range keys {
		fmt.Println("after sort", k, m[k])
	}
}
