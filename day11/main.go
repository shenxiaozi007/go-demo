package main

import (
    "fmt"
    "log"
    "math/rand"
    "strconv"
    "sync"
    "time"
)

func main()  {
    wg := &sync.WaitGroup{}

    wg.Add(10)

    //创建个管道
    ch := make(chan int, 5)

    //多个读
    for i := 0; i < 10; i++ {
        go func(i int) {
            defer wg.Done()
            for  {
                time.Sleep(time.Second * 2)
                select {
                case data, ok := <-ch:
                    if !ok {
                        fmt.Printf("关闭=%s \r\n", i)
                        return
                    }
                    //输出
                    fmt.Printf("i=%s, data=%s \r\n", i, data)
                }
            }
        }(i)
    }

    //单个写
    for i := 0; i <= 20; i++ {
        ch <- i
    }

    fmt.Println(ch)
    close(ch)
    wg.Wait()
    //结束
    fmt.Println("完成")
}


func main2() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    // ...
    const Max = 100000
    const NumReceivers = 10
    const NumSenders = 1000

    wgReceivers := sync.WaitGroup{}
    wgReceivers.Add(NumReceivers)

    // ...
    dataCh := make(chan int)
    stopCh := make(chan struct{})
    // stopCh is an additional signal channel.
    // Its sender is the moderator goroutine shown
    // below, and its receivers are all senders
    // and receivers of dataCh.
    toStop := make(chan string, 1)
    // The channel toStop is used to notify the
    // moderator to close the additional signal
    // channel (stopCh). Its senders are any senders
    // and receivers of dataCh, and its receiver is
    // the moderator goroutine shown below.
    // It must be a buffered channel.

    var stoppedBy string

    // moderator
    go func() {
        stoppedBy = <-toStop
        close(stopCh)
    }()

    // senders
    for i := 0; i < NumSenders; i++ {
        go func(id string) {
            for {
                value := rand.Intn(Max)
                if value == 0 {
                    // Here, the try-send operation is
                    // to notify the moderator to close
                    // the additional signal channel.
                    select {
                    case toStop <- "sender#" + id:
                    default:
                    }
                    return
                }

                // The try-receive operation here is to
                // try to exit the sender goroutine as
                // early as possible. Try-receive and
                // try-send select blocks are specially
                // optimized by the standard Go
                // compiler, so they are very efficient.
                select {
                case <- stopCh:
                    return
                default:
                }

                // Even if stopCh is closed, the first
                // branch in this select block might be
                // still not selected for some loops
                // (and for ever in theory) if the send
                // to dataCh is also non-blocking. If
                // this is unacceptable, then the above
                // try-receive operation is essential.
                select {
                case <- stopCh:
                    return
                case dataCh <- value:
                }
            }
        }(strconv.Itoa(i))
    }

    // receivers
    for i := 0; i < NumReceivers; i++ {
        go func(id string) {
            defer wgReceivers.Done()

            for {
                // Same as the sender goroutine, the
                // try-receive operation here is to
                // try to exit the receiver goroutine
                // as early as possible.
                select {
                case <- stopCh:
                    return
                default:
                }

                // Even if stopCh is closed, the first
                // branch in this select block might be
                // still not selected for some loops
                // (and forever in theory) if the receive
                // from dataCh is also non-blocking. If
                // this is not acceptable, then the above
                // try-receive operation is essential.
                select {
                case <- stopCh:
                    return
                case value := <-dataCh:
                    if value == Max-1 {
                        // Here, the same trick is
                        // used to notify the moderator
                        // to close the additional
                        // signal channel.
                        select {
                        case toStop <- "receiver#" + id:
                        default:
                        }
                        return
                    }

                    log.Println(value)
                }
            }
        }(strconv.Itoa(i))
    }

    // ...
    wgReceivers.Wait()
    log.Println("stopped by", stoppedBy)
}

func main1() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    // ...
    const Max = 1000000
    const NumReceivers = 10
    const NumSenders = 1000
    const NumThirdParties = 15

    wgReceivers := sync.WaitGroup{}
    wgReceivers.Add(NumReceivers)

    // ...
    dataCh := make(chan int)     // will be closed
    middleCh := make(chan int)   // will never be closed
    closing := make(chan string) // signal channel
    closed := make(chan struct{})

    var stoppedBy string

    // The stop function can be called
    // multiple times safely.
    stop := func(by string) {
        select {
        case closing <- by:
            <-closed
        case <-closed:
        }
    }

    // the middle layer
    go func() {
        exit := func(v int, needSend bool) {
            close(closed)
            if needSend {
                dataCh <- v
            }
            close(dataCh)
        }

        for {
            select {
            case stoppedBy = <-closing:
                exit(0, false)
                return
            case v := <- middleCh:
                select {
                case stoppedBy = <-closing:
                    exit(v, true)
                    return
                case dataCh <- v:
                }
            }
        }
    }()

    // some third-party goroutines
    for i := 0; i < NumThirdParties; i++ {
        go func(id string) {
            r := 1 + rand.Intn(3)
            time.Sleep(time.Duration(r) * time.Second)
            stop("3rd-party#" + id)
        }(strconv.Itoa(i))
    }

    // senders
    for i := 0; i < NumSenders; i++ {
        go func(id string) {
            for {
                value := rand.Intn(Max)
                if value == 0 {
                    stop("sender#" + id)
                    return
                }

                select {
                case <- closed:
                    return
                default:
                }

                select {
                case <- closed:
                    return
                case middleCh <- value:
                }
            }
        }(strconv.Itoa(i))
    }

    // receivers
    for range [NumReceivers]struct{}{} {
        go func() {
            defer wgReceivers.Done()

            for value := range dataCh {
                log.Println(value)
            }
        }()
    }

    // ...
    wgReceivers.Wait()
    log.Println("stopped by", stoppedBy)
}