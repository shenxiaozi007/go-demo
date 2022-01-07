package main

import "fmt"

func incr(p *int) int {
    *p++
    return *p
}
func main1()  {
    v := 1
    incr(&v)
    fmt.Println(v)
}

var gvar int

func main2()  {
    var one int
    _ = one
    two := 2
    fmt.Println(two)
    var three int
    three = 3

    one = three

    func(unused string) {
        fmt.Println("fuck")
    } ("what?")
}
//变量 one、two 和 three 声明未使用。知识点：未使用变量。如果有未使用的变量代码将编译失败

type ConfigOne struct {
    Daemon string
}

func (c *ConfigOne) String() string {
    //return fmt.Sprintf("fuck,: %v", c)
    return "fuck"
}

func main3()  {
    c := &ConfigOne{}
    fmt.Println(c.String())
}

//参考答案及解析：运行时错误。如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。

func main4()  {
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