package main

import (
	"fmt"
	"time"
)

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
//但是如果不特别指定的话，Go 语言不能识别类型，所以会报错:
//cannot use nil as type string in return argument.

func main3() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)
	}

	for _, f := range x {
		fmt.Println(f)
	}
}

type User struct{}
type User1 User
type User2 = User

func (i User1) m1() {
	fmt.Println("m1")
}

func (i User) m2() {
	fmt.Println("m2")
}

func main4()  {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}

//参考答案及解析：能，输出m1 m2，第 2 行代码基于类型 User 创建了新类型 User1，
//第 3 行代码是创建了 User 的类型别名 User2，注意使用 = 定义类型别名。因为 User2 是别名，
//完全等价于 User，所以 User2 具有 User 所有的方法。但是 i1.m2() 是不能执行的，因为 User1 没有定义该方法。

func main5()  {
	ch := make(chan int, 100)

	// A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}

			fmt.Println("a: ", a)
		}
	}()

	//close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}
//参考答案及解析：程序抛异常。先定义下，第一个协程为 A 协程，第二个协程为 B 协程；当 A 协程还没起时，
//主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic，
//panic: send on closed channel。

func main6() {
	ch := make(chan bool, 1)
	Stop(ch)
}

//func Stop(stop <-chan bool)  {
func Stop(stop chan bool)  {
	close(stop)
}

//参考答案及解析：有方向的 channel 不可以被关闭。

type Param map[string]interface{}

type Show struct {
	*Param
}

func main()  {
	s := new(Show)
	//s.Param["test"] = "2"

	p := make(Param)
	p["day"] = "2"
	s.Param = &p

	tmp := *s.Param
	fmt.Println(tmp["data"])
}
//考答案及解析：存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引。