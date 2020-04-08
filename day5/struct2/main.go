package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type member struct {
	a int8
	b int8
	c int8
	d int8
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

	//使用键值对对结构体进行初始化时，
	//键对应结构体的字段，值对应该字段的初始值
	p4 := person{
		name: "小王子",
		city: "beijin",
		age:  12,
	}
	fmt.Printf("p4 = %#v\n", p4)

	//也可以对结构体指针进行键值对初始化，例如：
	p5 := &person{
		name: "小碗中22",
		city: "beijin2",
		age:  10,
	}
	fmt.Printf("p5 = %#v\n", p5)

	//当某些字段没有初始值的时候，该字段可以不写。此时，
	//没有指定初始值的字段的值就是该字段类型的零值。
	p6 := &person{
		city: "北京",
	}

	fmt.Printf("p6= %#v \n", p6)

	//初始化结构体的时候可以简写，
	//也就是初始化的时候不写键，直接写值
	p7 := &person{
		"lailailai",
		"北京22",
		21,
	}
	fmt.Printf("p7 = %#v \n", p7)
	// 使用这种格式初始化时，需要注意：

	// 必须初始化结构体的所有字段。
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 该方式不能和键值初始化方式混用。

	n := &member{
		1, 2, 3, 4,
	}

	fmt.Printf("a- %p \n", &(*n).a)
	fmt.Printf("a- %p \n", &n.b)
	fmt.Printf("a- %p \n", &n.c)
	fmt.Printf("a- %p \n", &n.d)

}

func test(p person) {
	p.age = 10
}

func test2(p *person) {
	p.age = 20
}
