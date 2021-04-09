package main

import (
	"fmt"
	"sync"
)

func main() {

	wait := sync.WaitGroup{}
	wait.Add(2)
	intChan := make(chan int)

	//奇数
	go one(intChan, wait)
	//偶數
	go two(intChan, wait)
	intChan <- 0
	//wait.Wait()
}

func one(intChan chan int, waitInt sync.WaitGroup) {
	defer waitInt.Done()
	//str := "12345678"
	str := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for {
		oneInt, ok := <-intChan
		if !ok {
			return
		}
		if oneInt%2 == 0 {
			if oneInt == len(str) {
				close(intChan)
				fmt.Println("推出3")
				return
			}
			print(str[oneInt])
			intChan <- oneInt + 1
		}else {
			intChan <- oneInt
		}

	}
}

func two(intChan chan int, waitInt sync.WaitGroup) {
	defer waitInt.Done()
	//str := "abcdefgi"
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for {
		twoInt, ok := <-intChan
		if !ok {
			fmt.Println("推出2")
			return
		}
		if twoInt%2 == 1 {
			if twoInt >= len(str) {
				return
			}
			print(str[twoInt])
			intChan <- twoInt + 1
		} else {
			intChan <- twoInt
		}
	}
}

func test1(a []int) []int {
	test := []int{1, 2, 3}

	a = append(a, test...)
	fmt.Println(a)
	return a
}


