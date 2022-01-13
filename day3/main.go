package main

import "fmt"

type T struct {
	n int
}
func main1() {
	m := make(map[int]T)
	tmp := m[0]
	tmp.n = 1
	fmt.Println(m[0].n)
}

type X struct {}

func (x *X) test() {
	println(x)
}

func main() {
	var a *X
	a.test()
	var b = X{}
	// X{}.test()
	b.test()
}

//X{} 是不可寻址的，不能直接调用方法。知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。