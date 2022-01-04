package main

import (
	"fmt"
	"time"
)

//worker pool（goroutine池）
//在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
//一个简易的work pool示例代码如下：

func outputResults(out chan int, test chan int ) {
	for i := 0 ; i< 5 ; i++  {
		select {
		case test := <-out:
			fmt.Printf("测试 %v \n", test)
		}
		select {
			case test<- i:
				fmt.Printf("测试out %v \n", test)
		}
	}

}

func setResults(id int, out chan int, in chan int) {

		//i := <-out
		i := 1
		fmt.Printf("worker:%d start job:%d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, i)
		in <- i * 2

/*	for x := range out  {

		//写入results
		in <- x * 2
	}*/
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	test1 := make(chan int, 10)


	//写入
	for i := 0; i< 5 ; i++  {
		jobs<- i
	}

	// 5个任务
	for i := 0; i< 5; i++{
		go setResults(i, jobs, results)
	}

	close(jobs)
	outputResults(results, test1)
	out1 := <-test1
	fmt.Printf("out1 %v \n", out1)

}