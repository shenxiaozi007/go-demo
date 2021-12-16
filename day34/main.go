package main

import "fmt"

func f1() (r int) {
    defer func() {
        r++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t++
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}

func main1() {
    fmt.Println(f1()) //1
    fmt.Println(f2()) //5
    fmt.Println(f3()) //1
}

type test struct {
    name string
}

func (t *test) close() {
    fmt.Println(t.name, " closed")
}
func main2() {
    ts := []test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        //defer ts.close()  // 会输出 三个c
        // ts := t  用一个值保存
        defer func(t test) {
            t.close()
        }(t)
    }
}

//
type Person struct {
    age int
}

func main3() {
    preson := &Person{28}

    //1
    defer fmt.Println(preson.age)

    //2
    defer func(p *Person) {
        fmt.Println(p.age)
    }(preson)

    //3
    defer func() {
        fmt.Println(preson.age)
    }()

    //preson.age = 29
    preson = &Person{29}
}

//29 29 28
//1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
//
//2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
//
//3.闭包引用，输出 29；
//
//又由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 29 28。

//29 28 28
//    参考答案及解析：29 28 28。这道题在第 19 天题目的基础上做了一点点小改动，前一题最后一行代码 person.age = 29 是修改引用对象的成员 age，这题最后一行代码 person = &Person{29} 是修改引用对象本身，来看看有什么区别。
//
//    1处.person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
//
//    2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
//
//    3处.闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29.
//
//    由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 28 28。

func f() {
    defer fmt.Println("D")
    fmt.Println("F")
}
func main4() {
    f()
    fmt.Println("M")
}

// 被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M

type S struct {
}

func f4(x interface{}) {

}

func b5(x *interface{}) {

}
func main5() {
    s := S{}
    //p := &s

    f4(s)
    f4(s)
    //b5(s)
    //b5(p)
    //参考答案及解析：BD。函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 *interface{}。
    //永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
}

type S1 struct {
    m string
}

func f5() *S1 {
    return &S1{m: "foo"}
}

func main() {
    p := f5()
    fmt.Println(p.m)
}
