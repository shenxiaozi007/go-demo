package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//实例化
	var p1 person
	p1.name = "nimei"
	p1.city = "北京"
	p1.age = 18
	test(p1)
	fmt.Printf("p1 = %v\n", p1)
	fmt.Printf("p1 = %#v\n", p1)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小丸子"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//创建指针类型结构体
	p2 := new(person)
	p2.age = 21
	p2.city = "test"
	p2.name = "fuck"
	test2(p2)
	fmt.Printf("p2 = %#v \n", p2)

	//取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T \n", p3)
	fmt.Printf("p3 = %#v\n", p3)
	(*p3).age = 10
	p3.name = "test"
	p3.city = "test"
	fmt.Printf("p3 = %#v \n", p3)
	//p3.name = "test"其实在底层是(*p3).name = "test"，这是Go语言帮我们实现的语法糖。
}

func test(p person) {
	p.age = 10
}

func test2(p *person) {
	p.age = 20
}
