package main

import "fmt"

const (
    a = iota
    b = iota
)

const (
    name = "name"
    c = iota
    d = iota
)
func main1()  {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}


type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func main2()  {
    fmt.Println(South)
}
//参考答案及解析：South。知识点：iota 的用法、类型的 String() 方法。
//
//根据 iota 的用法推断出 South 的值是 2；另外，如果类型定义了 String() 方法，当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印。

type Math struct {
    x, y int
}
var m = map[string]*Math{
    "foo": &Math{2, 3},
}

func main3() {
    m["foo"].x = 4 //不能直接赋值

    //---- 1.使用临时变量
    tmp := m["foo"]
    tmp.x = 4
    m["foo"] = tmp
    //----

    //2.修改数据结构

    fmt.Println(m["foo"].x)
}

//编译报错 cannot assign to struct field m[“foo”].x in map。错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。

func main4()  {
    //fmt.Println([...]int{1} == [2]int{1})
    //fmt.Println([]int{1} == []int{1})
}


//go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 […]int{1} 和 [2]int{1} 是两种不同的类型，不能比较；
//切片是不能比较的；


var p *int

func foo() (*int, error) {
    var i int = 5
    return &i, nil
}
func bar()  {
    fmt.Println(*p)
}
func main() {
    //p, err := foo()//新定义的局部变量会覆盖全局变量
    var err error
    p, err = foo()
    if err != nil {
        fmt.Println(err)
        return
    }
    bar()
    fmt.Println(*p)
}

//参考答案及解析：B。知识点：变量作用域。问题出在操作符:=，对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。对于本例来说，main() 函数里的 p 是新定义的变量，会遮住全局变量 p，导致执行到bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash。
//
//正确的做法是将 main() 函数修改为：