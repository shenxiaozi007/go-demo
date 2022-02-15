package main

import (
    "fmt"
    "time"
)

func main1() {
    var x interface{}
    var y interface{} = []int{1, 2}

    _ = x == y
    _ = x == y
    _ = y == y
}

var o = fmt.Print

func main2() {
    c := make(chan int, 1)
    fmt.Println([3]struct{}{})
    for range [3]struct{}{} {
        select {
        case <-c:
            o(2)
            c = nil
        case c <- 1:
            o(3)
        default:
            o(1)
        }
    }
}

func test1(ch chan string) {
    time.Sleep(time.Second * 5)
    ch <- "test1"
}

func test2(ch chan string) {
    time.Sleep(time.Second * 2)
    ch <- "test2"
}

func main3() {
    // 2个管道
    output1 := make(chan string)
    output2 := make(chan string)

    // 跑2个子协程 写数据
    go test1(output1)
    go test2(output2)

    //c2 := <-output2
    //c1 := <-output1
    //fmt.Println(c2)
    //fmt.Println(c1)

    select {
    case s1 := <-output1:
        fmt.Println("s1=", s1)
    case s2 := <-output2:
        fmt.Println("s2=", s2)
    }
}

func main5() {
    //创建2个通道
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)

    go func() {
        int_chan <- 1
    }()

    go func() {
        string_chan <- "hello"
    }()
    //随机打印
    select {
    case value := <-int_chan:
        fmt.Println("int:", value)
    case value := <-string_chan:
        fmt.Println("string:", value)
    default:
        fmt.Println("default")
    }
    fmt.Println("main结束")

}

func main6() {
    //创建管道
    output1 := make(chan string, 10)
    //子协程写数据
    go write(output1)

    //取数据
    for s := range output1 {
        fmt.Println("res:", s)
        time.Sleep(time.Second)
    }

}

func write(ch chan string) {
    for {
        select {
        //写数据
        case ch <- "hello":
            fmt.Println("write hello")
        default:
            fmt.Println("channel full")
            //close(ch)
        }
        time.Sleep(time.Millisecond * 500)
    }
}

//超时执行
var resChan = make(chan int)

func main7() {
    go setData(resChan)
    select {
    case data := <-resChan:
        doData(data)
    case <-time.After(time.Second * 3):
        fmt.Println("request time out")
    }
}

func setData(data chan int) {
    time.Sleep(time.Second * 4)
    data <- 1
}

func doData(data int) {
    fmt.Println(data)
}

//退出操作

var shouldQuit = make(chan struct{})

func main8() {
    go closeChan()
    select {
    case v, ok := <-shouldQuit:
        fmt.Println(v, ok)
        closeUp()
        return
        /*default:
          fmt.Println("test")*/
    }
}

func closeChan() {
    fmt.Println("close")
    time.Sleep(time.Second * 2)
    close(shouldQuit)
}
func closeUp() {
    fmt.Println("close new")
}

//判断channel是否阻塞
func main9() {
    //判断channel是否阻塞
    ch := make(chan int, 5)

    //

    data := 0
    for  {
        select {
        case ch <- data:
            fmt.Println("test")
        default:
            //
            fmt.Println("fuck")
            time.Sleep(time.Second * 5)
        }
    }
}

type T struct {
    x int
    y *int
}

func main10() {
    i := 20
    t := T{10, &i}
    p := &t.x

    *p++
    *p--
    t.y = p

    fmt.Println(*t.y)
}

type TT struct {
    //num int
}

func (tt *TT) foo() {
    fmt.Println("fuck foo1")

    fmt.Println()
}

func (tt TT) bar() {
    fmt.Println("fuck bar1")
}

type SS struct {
    *TT
}

func main11()  {
    s := SS{&TT{}}
    fmt.Printf("%#v", s)
    _ = s.foo
    s.foo()
    s.bar()
    //_ = s.bar
}