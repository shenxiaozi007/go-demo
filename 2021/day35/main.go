package main

import (
    "fmt"
)

func main1() {
    //var x string = nil  //golang 的字符串类型是不能赋值 nil 的
    var x string = ""

    //if x == nil { //也不能跟 nil 比较。
    if x == "" {
        x = "default"
    }

    fmt.Println(x)
}

var a bool = true

func main2() {
    defer func() {
        fmt.Println("1")
    }()

    if a == true {
        fmt.Println("2")
        return
    }

    defer func() {
        fmt.Println("3")
    }()
}

//只有2, 1 defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法

func main3() {
    s1 := []int{1, 2, 3}
    s2 := s1[1:]
    s2[1] = 4
    fmt.Println(s1)
    s2 = append(s2, 5, 6,7)
    fmt.Println(s1)
}

//  1, 2, 4
// 1, 2, 4
//golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，这会导致 s2[1] = 4 语句影响 s1。
//
//而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1。

func main4() {
    if a := 1; false {

    } else if b := 2; false {

    } else {
        println(a, b)
    }
}
//  1, 2 知识点：代码块和变量作用域。

func main5() {
    m := map[int]string{0:"zero", 1:"one"}
    for k, v := range m {
        fmt.Println(k, v)
    }
}

//0 zero
//1 one
//或者
//1 one
//0 zero
//map 的输出是无序的。

func main6() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
}

func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)

    return ret
}


//main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：calc(“10”,a,b)，输出：10 1 2 3，得到值 3，因为
//defer 定义的函数是延迟函数，故 calc(“1”,1,3) 会被延迟执行；
//
//程序执行到第五行的时候，同样先执行 calc(“20”,a,b) 输出：20 0 2 2 得到值 2，同样将 calc(“2”,0,2) 延迟执行；
//
//程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc(“2”,0,2)，calc(“1”,1,3)，则就依次输出：2 0 2 2，1 1 3 4。




type Myint int

func (i Myint) PrintInt() {
    fmt.Println(i)
}

func main7()  {
    var i Myint = 1
    i.PrintInt()
}

//func (i int) PrintInt ()  {
//    fmt.Println(i)
//}
//
//func main() {
//    var i int = 1
//    i.PrintInt()
//}
//参考答案及解析：B。基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。

type People interface {
    Speak(string) string
}

type Student struct {

}

func (Stu *Student) Speak1(think string) (talk string)  {
    if think == "speak" {
        talk = "speak"
    } else {
        talk = "hi"
    }
    return
}

func (stu1 Student) Speak(think string) (talk string) {
    think = "11"
    return
}

func main8()  {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}

//编译错误 Student does not implement People (Speak method has pointer receiver)，值类型 Student 没有实现接口的 Speak() 方法，而是指针类型 *Student 实现该方法。
const (
    a1 = iota
    b1 = iota
)

const (
    name = "name"
    c1 = iota
    d1 = iota
)
func main10()  {
    fmt.Println(a1)
    fmt.Println(b1)
    fmt.Println(c1)
    fmt.Println(d1)
}

//参考答案及解析：0 1 1 2。知识点：iota 的用法。
//iota可以理解成const块中的行索引
//iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。
//
//iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。
//
//推荐阅读：
//
//https://studygolang.com/articles/2192


type People1 interface {
    Show()
}
type Student1 struct {

}

func (stu *Student1) Show() {

}
func main()  {
    var s *Student1
    if s == nil {
        fmt.Println("s is nil")
    } else {
        fmt.Println("s is not nil")
    }
    var p People1 = s
    if p == nil {
        fmt.Println("p is nil")
    } else {
        fmt.Println("p is not nil")
    }
}
//s is nil 和 p is not nil。这道题会不会有点诧异，我们分配给变量 p 的值明明是 nil，然而 p 却不是 nil。记住一点，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。上面的代码，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立

