package main

import "fmt"

func main1()  {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

func change(s ...int) {
	s = append(s, 3)
	fmt.Println(s)
}
//知识点：可变函数、append()操作。Go 提供的语法糖…，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；第二次调用change() 函数时，
//使用了操作符[i,j]获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。

func main2() {
	var a  = []int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

//[1, 12, 13, 4, 5]
//[1, 12, 13, 4, 5]
//这道题是 第30天 的第二题的一个解决办法，这的 a 是一个切片，那切片是怎么实现的呢？切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素

type Foo struct {
	bar string
}

func main3() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		//s2[i] = &value

		s2[i] = &s1[i]
		fmt.Println(&value)
	}


	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

//参考答案及解析：s2 的输出结果错误。s2 的输出是 &{C} &{C} &{C}，在 第 30 天 的答案解析第二题，我们提到过，for range 使用短变量声明(:=)的形式迭代变量时，变量 i、value 在每次循环体中都会被重用，而不是重新声明。所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。因此，s2 输出的时候显示出了三个 &{c}。

func main()  {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

//参考答案及解析：C。for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。