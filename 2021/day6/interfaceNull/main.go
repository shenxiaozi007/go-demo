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

	//类型断言
	//想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：x.(T)
	// x：表示类型为interface{}的变量
	// T：表示断言x可能是的类型。
	var xInterface interface{}
	x1 = "string"

	v, ok := xInterface.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	switchType(x1)
	switchType(x)
	switchType(test)

}

//空接口的应用
//使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type: %T value:%v", a, a)
}

func switchType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x type is a %v \n", v)
	case int:
		fmt.Printf("x type is a %v \n", v)
	case bool:
		fmt.Printf("x type is a %v \n", v)
	default:
		fmt.Printf("其他类型 %v\n", v)
	}
}
