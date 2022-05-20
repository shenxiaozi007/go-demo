package main

import (
    "context"
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type Ban struct {
    visitIPs map[string]time.Time
    rw sync.RWMutex
}

func NewBan(ctx context.Context) *Ban {
    banData := &Ban{visitIPs: make(map[string]time.Time)}

    go func() {
        //定时器
        timer := time.NewTicker(time.Second * 3)
        for  {
            select {
                //
                case <-timer.C:
                    fmt.Println("time start")
                    banData.rw.Lock()
                    for k, v := range banData.visitIPs{
                        if (time.Now().Sub(v) > time.Second * 3) {
                            fmt.Println(k, time.Now().Sub(v))
                            delete(banData.visitIPs, k)
                        }
                    }
                    banData.rw.Unlock()
                    //timer.Reset(time.Second * 3)
                case <-ctx.Done():
                    fmt.Println("time down")
                    return

            }
        }
    }()
    return banData
}
func (o *Ban) visit(ip string) bool {
    //加锁
    defer o.rw.Unlock()
    o.rw.Lock()
    if _, ok := o.visitIPs[ip]; ok {
        return true
    }
    o.visitIPs[ip] = time.Now()
    return false
}

//实现阻塞读且并发安全的map
func main1() {
    ctx, cancel := context.WithCancel(context.Background())
    defer time.Sleep(time.Second * 3)
    defer cancel()
    ban := NewBan(ctx)
    wg := sync.WaitGroup{}
    wg.Add(100 * 10)
    var success int64
    success = 0
    for i := 0; i < 100; i++ {
        for j := 0; j < 10; j++ {
            go func(j int) {
                defer wg.Done()
                ip := fmt.Sprintf("192.168.1.%d", j)
                if !ban.visit(ip) {
                    atomic.AddInt64(&success, 1)
                }
            }(j)
        }
        time.Sleep(time.Millisecond * 100)
    }
    wg.Wait()
    fmt.Println("success:", success)

}

//定时与 panic 恢复
func main()  {
    go func() {
        timer := time.NewTicker(time.Second * 2)
        for {

            select {
                case <-timer.C:
                    go func() {
                        defer func() {
                            if err := recover(); err != nil {
                                fmt.Println("err", err)
                            }
                        }()
                        //调用
                        proc()
                    }()

            }
        }
    }()

    select {}
}

func proc()  {
    panic("ok")
}