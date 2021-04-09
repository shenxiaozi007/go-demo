package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//一直循环
func main1() {
	wg.Add(1)
	go func() {
		for {
			fmt.Println("worker")
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("over")
}

//------怎么优雅的关闭
//全局变量方法
var exit bool

func main2() {

	wg.Add(1)

	go func() {
		for {
			fmt.Println("worker")
			time.Sleep(1 * time.Second)

			if exit {
				break
			}
		}
		wg.Done()
	}()

	time.Sleep(5 * time.Second)
	exit = true
	wg.Wait()
	fmt.Println("over")
}

//通过通道关闭

func main3() {

	var exitChan chan bool

	exitChan = make(chan bool)
	go func() {
		for {
			select {
			case <-exitChan:
				goto Loop
			default:
				fmt.Println("worker")
				time.Sleep(1 * time.Second)

			}
		}
	Loop:
	}()

	time.Sleep(5 * time.Second)
	exitChan <- false
	fmt.Println("over")
}

//通过context来取消

func main() {

}
