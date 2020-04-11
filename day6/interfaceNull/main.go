package main

import "fmt"

func main() {
	// 空接口的定义
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。
	var x interface{}

	x1 := "Hello 空接口"
	x = x1
	fmt.Println(x)

	x2 := 2
	x = x2
	fmt.Println(x)

	show(12)
	show("32")

	//空接口作为map的值
	//使用空接口实现可以保存任意值的字典。
	test := func() {
		mapInfo := make(map[string]interface{})
		mapInfo["int"] = 11
		mapInfo["string"] = "test"
		mapInfo["bool"] = false
		fmt.Println(mapInfo)
	}
	test()

	//
}

//空接口的应用
//使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type: %T value:%v", a, a)
}
