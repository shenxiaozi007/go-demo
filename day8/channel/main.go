package main

import (
	"fmt"
)

// channel是一种类型，一种引用类型。声明通道类型的格式如下：
// var 变量 chan 元素类型

var ch1 chan int // 声明一个传递整型的通道
var ch2 chan []string // 声明一个传递string切片的通道

func recv(c chan int) {
	ret := <- c
	fmt.Printf("接受值-%v", ret)
}

// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

//<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

//of the reader
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 = make(chan int)

	go recv(ch1)       // 从ch中接收值，忽略结果
	// channel操作
	//通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送和接收都使用<-符号。
	ch1 <- 10

	fmt.Println("发送成功")

	// 创建一个容量为1的有缓冲区通道

	ch2 := make(chan int, 2)
	ch2 <- 20
	ret := <- ch2
	fmt.Println(ret)

	//for range从通道循环取值

	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch4 <- i
		}
		close(ch4)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch3 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	for i := range ch4 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	//单向通道
	//有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
	//Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：

	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)


}
