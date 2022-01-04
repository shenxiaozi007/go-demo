package main

import "fmt"

func foo() (int, string) {
	return 10, "test"
}

func main() {
	x, _ := foo()
	_, y := foo()

	fmt.Println(x, y)

	//多个声明
	var (
		a string
		b int
		c bool
		d string
	)

	a = "32"
	b = 2
	c = false
	d = "32"

	fmt.Println(a, b, c, d)

	//短变量声明
	f := 3232
	fmt.Println(f)
}
