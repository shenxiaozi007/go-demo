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

func main2()  {
    fmt.Println(x,y,z,k,p)
}



//func GetValue() int {
func GetValue() interface{} {
    return 1
}

func main3() {
    i := GetValue()
    switch i.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    }
}
//编译失败。考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，
//type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。看下关于接口的文章。

func hello() []string {
    return nil
}

func main4()  {
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}

type person struct {
    name string
}

func main5() {
    var m map[person]int

    p := person{"mike"}

    var m1 map[string]interface{}
    fmt.Println(m[p])
    fmt.Println(m1["1"])
}
//打印一个 map 中不存在的值时，返回元素类型的零值。这个例子中，
//m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0。

func