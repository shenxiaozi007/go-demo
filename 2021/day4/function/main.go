package main

import "fmt"

//定义全局变量
var num int64 = 10

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

//多返回值
// Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名
// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calcTwo(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//返回值补充
// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func calcThree(x int) []int {
	if x == 0 {
		return nil
	}
	data := []int{}
	data = append(data, x)
	return data
}

//语句块定义的变量 只能在语句内使用
func testLocalVal(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100
		fmt.Println(z)
	}
	//fmt.Println(z) //此处无法使用变量z
}
