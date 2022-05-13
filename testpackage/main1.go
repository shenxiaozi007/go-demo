package main

import (
	"fmt"
	"sync"
)

func main1() {
	wait := sync.WaitGroup{}

	wait.Add(2)

	umInt := make(chan int, 1)
	umStr := make(chan int, 1)
	go putString(umInt, umStr, wait)
	go putMumber(umInt, umStr, wait)

	wait.Wait()
	fmt.Println(3232)
}

func putMumber(umInt chan int, umStr chan int, group sync.WaitGroup) {
	defer close(umStr)
	defer group.Done()

	intStr := "12345678"
	umInt <- 0
	for {
		for value := range umInt {
			if value%2 != 1 {
				if value+2 > len(intStr) {
					umStr <- value
					return
				}
				print(intStr[value : value+2])
				umStr <- value + 1
			} else {
				umStr <- value
			}
		}
	}
}

func putString(umInt chan int, umStr chan int, group sync.WaitGroup) {
	defer close(umInt)
	defer group.Done()

	stringStr := "abcdefgi"
	for {
		for value := range umStr {
			if value%2 == 1 {
				if value+2 > len(stringStr) {
					umInt <- value
					return
				}
				print(stringStr[value-1 : value+1])
				umInt <- value + 1
			} else {
				umInt <- value
			}
		}
	}
}