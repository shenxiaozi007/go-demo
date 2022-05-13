package main

import (
    "fmt"
    "time"
)

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
func main5() {
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

func main6() {
    v := []int{1, 2, 3}
    for i := range v {
        v = append(v, i)
    }
    fmt.Println(v)
}
//参考答案及解析：不会出现死循环，能正常结束。
//循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。

func main7()  {
    var m = [...]int{1, 2, 3}
    for i, v := range m {
        go func() {
           fmt.Println(i, v)
        }()
        //------1.使用函数传递
        go func(i int, v int) {
            fmt.Println(i, v)
        }(i, v)

        //------2 临时变量
        i_tmp := i  // 这里的 := 会重新声明变量，而不是重用
        v_tmp := v
        go func() {
            fmt.Println(i_tmp, v_tmp)
        }()

    }
    time.Sleep(time.Second * 3)
}
//for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。
//
//各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，而不是各个goroutine启动时的i, v值。可以理解为闭包引用，使用的是上下文环境的值。
//2,3
//2,3
//2,3

func f(n int) (r int) {
    defer func() {
        r += n
        recover()
    }()

    var f func()

    defer f()
    f = func() {
        r += 2
    }
    return n + 1
}

func main8()  {
    fmt.Println(f(3))
}

//参考答案及解析：7。根据 5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！ 提到的“三步拆解法”，
//第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。

func main()  {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range &a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }

    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}

//[1, 2, 3, 4, 5]
//[1,12,13,4,5]
//range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a。就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的

//r =  [1 12 13 4 5]
//a =  [1 12 13 4 5]
//修复代码中，使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。


