package main

import (
	"fmt"
	"sync"
)

var s []int
var wg sync.WaitGroup
var lock sync.Mutex
var test chan int

func appendValue(test chan int) {
	defer wg.Done()
	lock.Lock()
	s = append(s, <-test)
	lock.Unlock()
}

func main() {

	//map 非线程安全

	m := make(map[int]int)

	go func() {
		for i := 0 ; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()


	//切片是非线程安全的
	test = make(chan int, 1)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		test <- i
		go appendValue(test)
	}

	for k, v := range s {
		fmt.Println( k, ":" ,v)
	}


	wg.Wait()
}


