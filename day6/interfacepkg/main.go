package main

import "fmt"

type dog struct{}
type cat struct{}

//Sayer is a interface
type Sayer interface {
	say()
}

//Mover is a interface
type Mover interface {
	move(run string)
}

//Mover2 is a interface
type Mover2 interface {
	move2()
}

//接口嵌套

func main() {
	var x Sayer
	dogS := dog{}
	catS := cat{}
	x = dogS
	dogS.say()
	catS.say()
	x.say()

	//值接收者实现接口
	tets := func() {
		var m Mover
		wancai := dog{} //wangcai is dog类型
		m = wancai
		m.move("wancai")
		fugui := &dog{} //fugui is *dog类型
		m = fugui
		m.move("fugui")

		//从上面的代码中我们可以发现，使用值接收者实现接口之后，
		//不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
		//因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui

	}
	tets()

	//指针接收者实现接口
	test2 := func() {
		var m Mover2
		// wangcai2 := dog{}
		//m = wangcai2 //m不可以接收dog类型
		wangcai2 := &dog{}
		m = wangcai2
		m.move2()
		//此时实现mover2接口的是*dog类型，所以不能给m传入dog类型的wangcai2，此时x只能存储*dog类型的值。
	}
	test2()
}

// 实现了say的方法
func (d dog) say() {
	fmt.Println("fuck")
}

func (d dog) move(run string) {
	fmt.Println(run + "赶快跑")
}

func (d *dog) move2() {
	fmt.Println("指针跑")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
