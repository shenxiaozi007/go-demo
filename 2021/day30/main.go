package main

import "fmt"

var c = make(chan int)
var a int

func f() {
    a = 1
    f := <-c
    fmt.Println(f)
}

func main()  {
    go f()
    c <- 0
    print(a)
}
