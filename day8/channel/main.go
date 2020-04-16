package main

// channel是一种类型，一种引用类型。声明通道类型的格式如下：
// var 变量 chan 元素类型

var ch1 chan int // 声明一个传递整型的通道
var ch2 chan []string // 声明一个传递string切片的通道
func main() {
	ch1 = make(chan int)

	// channel操作
	//通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送和接收都使用<-符号。

	ch1 <- 11 // 把10发送到ch中

	<-ch1       // 从ch中接收值，忽略结果
}
