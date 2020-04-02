package main

import "fmt"

func main() {
	x := 10
	y := 20
	sum := intSum(x, y)
	sum2 := intSumTwo(x, y)
	fmt.Println(sum, sum2)
	sayHello()
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
