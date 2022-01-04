package main

import "fmt"

const i1 = 100
var j = 123

func main1() {
	//fmt.Println(&i1, i1)
	fmt.Println(i1, i1)
	//编译报错cannot take the address of i。知识点：常量。常量不同于变量的在运行期分配内存，
	//常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
	fmt.Println(&j, j)
}

func GetValue(m map[int]string, id int) (string, bool)  {
	if _, exist := m[id]; exist {
		return "exist", true
	}
	//return nil, false
	return "", false
}

func main2()  {
	intmap := map[int]string {
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}

//不能通过编译。知识点：函数返回值类型。nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
//但是如果不特别指定的话，Go 语言不能识别类型，所以会报错:cannot use nil as type string in return argument.

func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)
	}

	for _, f := range x {
		fmt.Println(f)
	}
}


