package main

import (
    "context"
    "fmt"
    "runtime/debug"
    "time"
)

func main1() {
    var a [1]int
    c := a[:]
    println(c)
}

func main2()  {
    slice := make([]string, 2, 4)
    example(slice, "hello", 10)
}

func example(slice []string, str string, i int) {
    debug.PrintStack()
}

func deadLineTest(ctx context.Context)  {
    go func() {
        select {
        case <-time.After(time.Second * 5):
            break
        }
    }()
}
func main() {
    d := time.Now().Add(50 * time.Millisecond)

    ctx, cancel := context.WithDeadline(context.Background(), d)
    fmt.Println(ctx, cancel)

}