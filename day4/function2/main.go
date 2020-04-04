package main

import (
	"errors"
	"fmt"
)

//函数类型变量
type calculation func(int, int) int

func main() {
	var c calculation
	c = add

	fmt.Printf("type of f:%T\n", c)
	//
	fmt.Println(c(1, 2))

	f := sub

	fmt.Printf("type of f:%T\n", f)
	fmt.Println(f(2, 1))

	//函数作为参数
	g := calc(1, 2222, add)
	fmt.Println(g)

	//函数作为返回值
	h, err := do("-")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		res := h(1, 342)
		fmt.Println(res)
	}
	//匿名函数
	anonymous()

	//闭包
	var b = adder()
	fmt.Printf("type of f:%T", b)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//匿名函数
// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:
func anonymous() {
	//将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}

	add(10, 20)

	//自执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(20, 40)

}

//闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 首先我们来看一个例子
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
