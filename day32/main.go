package main

import "fmt"


//-----------
//新类型 MyInt1
type MyInt1 int

//别名
type MyInt2 = int


func main1()  {
    var i int = 0
    //Cannot use 'i' (type int) as type MyInt1
    //var i1 MyInt1 = i
    var i1 MyInt1 = MyInt1(i)
    var i2 MyInt2 = i

    fmt.Println(i1, i2)
}
//参考答案及解析：编译不通过，cannot use i (type int) as type MyInt1 in assignment。
//
//这道题考的是类型别名与类型定义的区别。
//
//第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。所以，
//第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
//----------

const (
    x = iota
    _
    y
    z = "zz"
    k
    p = iota
)

func main()  {
    fmt.Println(x,y,z,k,p)
}
