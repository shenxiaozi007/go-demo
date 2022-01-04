package main

import "fmt"

func main() {

	var a = [3]int{1, 2, 3}
	var b = []int{1, 2, 3}
	fmt.Println(a, b)
	fmt.Printf("a:%T b:%T", a, b)
	//从数组得到切片
	//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
	x := [...]string{"string", "1", "2", "3", "4", "5", "6"}
	y := x[1:4]
	fmt.Printf("切片, %d\n", len(y))
	fmt.Printf("切片, %d\n", cap(y))
	//append()方法为切片添加元素和扩容
	var sumSlice []int
	for i := 0; i < 10; i++ {
		sumSlice = append(sumSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", sumSlice, len(sumSlice), cap(sumSlice), sumSlice)
	}

	var citySlice []string

	citySlice = append(citySlice, "北京")

	p := []string{"test", "test2"}

	citySlice = append(citySlice, p...)

	fmt.Println(citySlice)
}
