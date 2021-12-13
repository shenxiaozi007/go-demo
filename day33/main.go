package main

import "fmt"

type PeoPle struct {
}

func (p *PeoPle) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}

func (p *PeoPle) ShowB() {
    fmt.Println("showB")
}


type Teacher struct {
    PeoPle
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}
func main1()  {
    t := Teacher{}
    t.ShowB()
    t.PeoPle.ShowB()
}

//知识点：结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，
//内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法


func incr(p *int) int {
    *p++
    return *p
}

func main2()  {
    p := 1
    incr(&p)
    fmt.Println(p)
}

//指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。

func add(args ...int) int {
    sum := 0

    for _, arg := range args {
        sum += arg
    }

    return sum
}

func main3() {
    add([]int{1, 2}...)
    add(1,2)
    add(1,2,7)
}

//ABD。知识点：可变函数。
func main4() {
    var s1 []int
    var s2 = []int{}
    if s2 == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("no nil")
    }

    if s1 == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("no nil")
    }
}

//知识点：nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main5()  {
    c := Work{i : 3}
    var a A = c
    var b B = c

    fmt.Println(a.ShowA())
    fmt.Println(b.ShowB())
}
//知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA()

func main() {
    s := [3]int{1, 2, 3}
    a := s[:0]
    b := s[:2]
    c := s[1:2:cap(s)]
}