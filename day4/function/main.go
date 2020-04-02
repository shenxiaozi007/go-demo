package main

import "fmt"

func main() {
	x := 10
	y := 20
	sum := intSum(x, y)
	sum2 := intSumTwo(x, y)
	fmt.Println(sum, sum2)
	sayHello()

	sumThree := intSumThree(1, 2, 3, 5, 6)
	fmt.Println(sumThree)
}

//func 函数名(参数)(返回值){
// 函数体
// }
func intSum(x int, y int) int {
	return x + y
}

//没有返回值的
func sayHello() {
	fmt.Println("hello 沙河")
}

//类型简写
func intSumTwo(x, y int) int {
	return x + y
}

// 变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
// 注意：可变参数通常要作为函数的最后一个参数。而且为一个切片
func intSumThree(x ...int) int {
	fmt.Println(x) //切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
	// 本质上，函数的可变参数是通过切片来实现的。
}
