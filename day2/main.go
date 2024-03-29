package main

import (
    "encoding/json"
    "fmt"
)

func incr(p *int) int {
	*p++
	return *p
}
func main1() {
	v := 1
	incr(&v)
	fmt.Println(v)
}

var gvar int

func main2() {
	var one int
	_ = one
	two := 2
	fmt.Println(two)
	var three int
	three = 3

	one = three

	func(unused string) {
		fmt.Println("fuck")
	}("what?")
}

//变量 one、two 和 three 声明未使用。知识点：未使用变量。如果有未使用的变量代码将编译失败

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	//return fmt.Sprintf("fuck,: %v", c)
	return "fuck"
}

func main3() {
	c := &ConfigOne{}
	fmt.Println(c.String())
}

//参考答案及解析：运行时错误。如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。

func main4() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r)
}

//参考答案及解析：[1 2 3 4 5]。a 在 for range 过程中增加了两个元素
//，len 由 5 增加到 7，但 for range 时会使用 a 的副本 a’ 参与循环，副本的 len 依旧是 5，因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素。

func main5() {
	//var x = nil
	var x interface{} = nil
	_ = x
}

//nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。
//如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误。

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func main6() {
	var data info
	var err error
	//data.result, err := work()
	data.result, err = work()

	if err != nil {
		fmt.Println()
	}
}

//不能使用短变量声明设置结构体字段值，修复代码：

func main7() {
	const x = 123
	const y = 1.23
	fmt.Println(y)
}

//编译可以通过。知识点：常量。常量是一个简单值的标识符，
//在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func main8() {
	fmt.Printf("%T %v \n", y, y)
	fmt.Printf("%T %v \n", z, z)

}

func main9() {
	//var x string = nil
	var x string = ""
	if x == "" {
		//if x = nil {
		x = "default"
	}
}

func main10() {
	var ch chan int

	//ch = make(chan int, 10)

	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}

}

type Peopleday3 struct {
    //name string `json:"name"`
    Name string `json:"name"`
}

func main11() {
    js := `{
        "name" : "seekload"
    }`
    var p Peopleday3
    err := json.Unmarshal([]byte(js), &p)

    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println(p)
}

type T struct {
    ls []int
}

func foo(t T) {
    t.ls[0] = 100
}

func main12() {
    var t = T{
        ls : []int{1, 2, 3},
    }

    foo(t)
    fmt.Println(t.ls[0])

}
//参考答案及解析：B。调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针

func main() {
    isMatch := func(i int) bool {
        switch (i) {
        case 1:
            fallthrough
        case 2:
            return true
        }
        return false
    }
    fmt.Println(isMatch(1))
    fmt.Println(isMatch(2))

    isMatch2 := func(i int) bool {
        switch (i) {
        case 1,2:
            return true
        }
        return false
    }
    fmt.Println(isMatch2(1))
    fmt.Println(isMatch2(2))
}
