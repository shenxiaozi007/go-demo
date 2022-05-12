package main

import (
    "fmt"
    "time"
)

func test1(i *int) {

}

func main() {
    i := 5

    t := test1

    t(&i)
    //p := &Prople{}
    //p.String11()
    //test2()

    //--- 2
    project := new(Project)
    project.Main()
}

type Prople struct {
    Name string
}

func (p *Prople) String11() string {
    return fmt.Sprintf("print: %v", p)
}

func test2() {
    ch := make(chan int, 1000)
    //wg := &sync.WaitGroup{}
    //wg.Add(1)
    go func() {
        //defer wg.Done()
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)
    }()

    go func() {
        for {
            a, ok := <-ch
            if !ok {
                fmt.Println("close")
                return
            }
            fmt.Println("a: ", a)
        }
    }()

    //wg.Wait()
    fmt.Println("ok")
    time.Sleep(time.Second * 10)
}

type Project struct{}

func (p *Project) deferError() {
    if err := recover(); err != nil {
        fmt.Println("recover11: ", err)
    }
}

func (p *Project) exec(msgchan chan interface{}) {
    defer p.deferError()
    for msg := range msgchan {
        m := msg.(int)
        fmt.Println("msg: ", m)
    }
}

func (p *Project) run(msgchan chan interface{}) {
    for {
        go p.exec(msgchan)
        time.Sleep(time.Second * 2)
    }
}

func (p *Project) Main() {
    a := make(chan interface{}, 100)
    go p.run(a)
    go func() {
        for {
            a <- "1"
            time.Sleep(time.Second)
        }
    }()
    time.Sleep(time.Second * 10)
}