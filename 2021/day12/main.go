package main

import (
	"fmt"
	"sync"
	"time"
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

func main2() {
	//切片是非线程安全的
	test = make(chan int, 1)
	for p := 0; p < 10000; p++ {
		wg.Add(1)
		test <- p
		go appendValue(test)
	}

	for k, v := range s {
		fmt.Println(k, ":", v)
	}

	wg.Wait()
}

func main() {
	//map 非线程安全
	m := make(map[int]int)

	go func() {
		lock.Lock()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
		lock.Unlock()
	}()

	go func() {
		lock.Lock()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
		lock.Unlock()
	}()
	time.Sleep(time.Second * 3)
}
