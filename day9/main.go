package main

import (
    "context"
    "fmt"
    "log"
    "time"
)

func main1() {
    //ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
    timeF := "2006-01-02"
    nowTime := time.Now()

    yesterdayTime := nowTime.AddDate(0, 0, 1)
    beginTime, _ := time.ParseInLocation(timeF, yesterdayTime.Format(timeF), time.Local)
    beginTime2, _ := time.Parse(timeF, yesterdayTime.Format(timeF))
    fmt.Printf("%v\n%v\n", beginTime, beginTime.Unix())
    fmt.Printf("%v\n%v\n", beginTime2, beginTime2.Unix())
    fmt.Println(time.Unix(beginTime.Unix(), 0))
    fmt.Println(time.Unix(beginTime2.Unix(), 0))
    fmt.Println(time.Since(yesterdayTime))

}
func longRunningCalculation(timeC int) chan string {
    res := make(chan string)
    go func() {
        time.Sleep(time.Second * time.Duration(timeC))
        res <- "Done"
    }()
    return res
}

func testCtx(parent context.Context) {
    select {
    case <-parent.Done():
        log.Println(parent.Err())
        break
    }

    fmt.Println("退出循环")
}

func jobHandle() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    go testCtx(ctx)
    defer cancel()
    select {
    case result := <-longRunningCalculation(5):
        fmt.Println(result)
    }
    return
}
func main() {
    jobHandle()
}
