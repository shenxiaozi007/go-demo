package main

import "fmt"

func main() {
	//每日练习 数组
	var a [3]int
	a[1] = 1
	a[2] = 2
	var b [3]string
	b[1] = "3232"
	fmt.Println(a)
	fmt.Printf("%T", a)
	//循环
	for i := 0 ; i < 3 ; i ++ {
		fmt.Println(a[i])
	}

	for val, key := range a {
		fmt.Println(val,key)
	}
	//
	for v, k := range b {
		fmt.Println(v,k)
	}
	//每日练习 切片 切片是引用类型

	var a1 []int
	//a1 = make([]int, 4)
	a2 := []string{}
	a3 := make([]bool, 2)

	//a1[3] = 1
	a1 = append(a1, 1)
	a2 = append(a2,"test")
	a3 = append(a3, true)
	fmt.Println(a1,a2,a3)

	a4 := a3
	a4[0] = true
	fmt.Println(a1,a2,a3,a4)


}
