package main

import "fmt"

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

}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
