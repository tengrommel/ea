package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var num int = 0

func add(paddr *int) {
	for i := 0; i < 10000; i++ {
		*paddr++
	}
	wg.Done()
}

func main() {
	var num = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add(&num)
	}
	wg.Wait()
	fmt.Println(num)
}
