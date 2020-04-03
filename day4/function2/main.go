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
