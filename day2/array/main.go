package main

import "fmt"

func main() {
	//数组定义
	//var 数组变量名 [元素数量]T
	var a = [3]int{1, 2, 3}
	fmt.Println(a)
	//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	var b = [...]string{"32", "32"}
	fmt.Println(b)
	//我们还可以使用指定索引值的方式来初始化数组，例如:
	d := [...]int{1: 3, 3: 5}
	fmt.Println(d)

	//数组遍历
	//第一种
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	//第二种
	for index, value := range b {
		fmt.Println(index, value)
	}
	//多维数组
	var f [3][2]int

	f = [3][2]int{
		[2]int{2, 3}, [2]int{3, 4}, [2]int{5, 6},
	}
	fmt.Println(f)

}
