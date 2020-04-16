package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func hello(i int) {
	fmt.Printf("hello goroutine -- %v \n", i)
}

func helloWg(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Printf("hello goroutine -- %v \n", i)
}
func main() {
	for i := 0; i < 10 ; i++  {
		wg.Add(1)// 启动一个goroutine就登记+1
		go helloWg(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("main goroutine done!")
}





