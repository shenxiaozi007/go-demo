package main

import (
    "fmt"
    "sync"
    "time"
)

type data struct {
    mut sync.Mutex
}

func (d data) test1(s string) {
    fmt.Printf("%p==||%p\n", &d, &d.mut)
    d.mut.Lock()
    defer d.mut.Unlock()

    for i := 0; i < 5; i++ {
        fmt.Println(s, i)
        time.Sleep(time.Second)
    }
}

func (d *data) test2(s string) {
    fmt.Printf("%p==||%p\n", d, &d.mut)
    d.mut.Lock()
    defer d.mut.Unlock()

    for i := 0; i < 5; i++ {
        fmt.Println(s, i)
        time.Sleep(time.Second)
    }
}


func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    //testInt := 1
    //nut := sync.Mutex{}
    d := data{}
    g := &d
    fmt.Printf("%p= \n", g)
    //var d data

    go func() {
        defer wg.Done()
        //d.test1("read")
        //g.test1("read2")
        d.test2("read")
        g.test2("read2")
    }()

    go func() {
        defer wg.Done()
        //d.test1("write")
        //g.test1("read2")
        d.test2("write")
        g.test2("read2")
    }()
    wg.Wait()
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

func main1() {
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