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


}
