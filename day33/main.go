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

func main() {
    add([]int{1, 2}...)
    add(1,2)
    add(1,2,7)
}

//ABD。知识点：可变函数。