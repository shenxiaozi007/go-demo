package main

import (
	"fmt"
	"go-demo/day5/pkg2"
	pkg "go-demo/day5/pkg2"
)

type Person struct {
	name string
	city string
	age  int
}

//Go语言的结构体没有构造函数，我们可以自己实现
func newPerson(name, city string, age int) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//方法和接收者
//Go语言中的方法（Method）是一种作用于特定类型变量的函数。
//这种特定类型变量叫做接收者（Receiver）。
//接收者的概念就类似于其他语言中的this或者 self。

// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
// 		函数体
// }

//传值
func (p Person) Dream() {
	fmt.Printf("这是一个方法%s \n", p.name)
}

//指针接受者
func (p *Person) SetAge(age int) {
	p.age = age
}

func main() {

	//调用构造函数
	p := newPerson("test", "meiguo", 12)
	fmt.Printf("%v \n", p)
	fmt.Println(p)

	//
	p.Dream()

	p.SetAge(32)
	fmt.Println(p)
	//相加
	fmt.Println(pkg2.Add(2, 3))
	fmt.Println(pkg.Add(1, 3))
}
