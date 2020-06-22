package main

import (
	"context"
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
				goto LOOP
			default:
				fmt.Println("worker")
				time.Sleep(1 * time.Second)

			}
		}
		LOOP:
	}()
	time.Sleep(5 * time.Second)
	exitChan <- false
	fmt.Println("over")

}

//官方的方案
func main4() {
	ctx, cancel := context.WithCancel(context.Background())

	//wg.Add(1)

	go func(ctx context.Context) {

		for  {
			select {
			case <-ctx.Done():
				goto LOOP
			default:
				fmt.Println("worker")
				time.Sleep(1 * time.Second)
			}

		}
	LOOP:
		//wg.Done()
	}(ctx)
	time.Sleep(5 * time.Second)

	cancel()//通知子goroutine结束

	//wg.Wait()

	fmt.Println("over")
}

//如果子goroutine里面又有另一个goroutine。
func main5() {
	cxt, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go func(ctx context.Context) {

		//子goroutine
		go func(ctx context.Context) {
			for  {
				select {
				case <-ctx.Done():
					fmt.Println("worker2 done")
					goto LOOP
				default:
					fmt.Println("worker2")
					time.Sleep(1 * time.Second)
				}
			}
			LOOP:
			wg.Done()
		}(ctx)

		for  {
			select {
			case <- ctx.Done():
					fmt.Println("worker1 done")
					goto LOOP
			default:
				fmt.Println("worker1")
				time.Sleep(1 * time.Second)
			}
		}
		LOOP:
		wg.Done()
	}(cxt)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("over")
}
