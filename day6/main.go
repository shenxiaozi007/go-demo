package main

import (
    "fmt"
    "sync"
    "time"
)

type data struct {
    mut *sync.Mutex
}

func (d *data) test1(s string) {
    fmt.Printf("%p==||%p\n", &d, &d.mut)
    d.mut.Lock()
    defer d.mut.Unlock()

    for i := 0; i < 5; i++ {
        fmt.Println(s, i)
        time.Sleep(time.Second)
    }
}

func main1() {

    var wg sync.WaitGroup
    wg.Add(2)
    //testInt := 1
    //nut := sync.Mutex{}
    mu := sync.Mutex{}
    d := &data{
        mut : &mu,
    }

    //g := &d
    //fmt.Printf("%p= \n", g)
    //var d data

    go func() {
        defer wg.Done()
        d.test1("read")
        //g.test1("read2")
        //d.test2("read")
        //g.test2("read2")
    }()

    go func() {
        defer wg.Done()
        d.test1("write")
        //g.test1("read2")
        //d.test2("write")
        //g.test2("read2")
    }()
    wg.Wait()
}

func (d *data) test2(s string) {
    //fmt.Printf("%p==||%p\n", d, &d.mut)
    d.mut.Lock()
    defer d.mut.Unlock()

    for i := 0; i < 5; i++ {
        fmt.Println(s, i)
        time.Sleep(time.Second)
    }
}
type Data struct {
    x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
    fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
    fmt.Printf("Pointer: %p\n", self)
}

func main2() {
    d := Data{}
    p := &d
    fmt.Printf("Data: %p\n", p)

    d.ValueTest()   // ValueTest(d)
    d.PointerTest() // PointerTest(&d)

    p.ValueTest()   // ValueTest(*p)
    p.PointerTest() // PointerTest(p)
}

//当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。
//可用实例 value 或 pointer 调用全部方法，编译器自动转换。

func main3() {
    var k = 1
    var s = []int{1, 2}
    k, s[k] = 0, 3
    fmt.Println(s[0] + s[1])

}

func main4() {
    var k = 9
    for k = range []int{} {}
    fmt.Println(k)
    for k = 0; k < 3; k++ {

    }

    fmt.Println(k)

    for k = range (*[3]int)(nil) {
        
    }

    fmt.Println(k)
}

func main5() {
    //预定义变量可以被覆盖
    nil := 123
    fmt.Println(nil)

    //var _ map[string]int = nil
}

func F(n int) func() int {
    return func() int {
        n++
        return n
    }
}

func main6() {
    f := F(5)

    defer func() {
        fmt.Println(f())
    }()

    defer fmt.Println(f())

    i := f()
    fmt.Println(i)
}

