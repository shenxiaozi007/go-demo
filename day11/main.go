package main

import (
    "fmt"
    "log"
    "math/rand"
    "strconv"
    "sync"
    "time"
)

func main3() {
    wg := &sync.WaitGroup{}

    wg.Add(10)

    //创建个管道
    ch := make(chan int, 5)

    //多个读
    for i := 0; i < 10; i++ {
        go func(i int) {
            defer wg.Done()
            for {
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
                case <-stopCh:
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
                case <-stopCh:
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
                case <-stopCh:
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
                case <-stopCh:
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
            case v := <-middleCh:
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
                case <-closed:
                    return
                default:
                }

                select {
                case <-closed:
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

//多个接收者，一个发送者，发送者关闭channel表示‘没有值可以发送’
//这是一个非常简单的情况，仅仅是让发送者在不想发送数据时候关闭channel。
func main4() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    //最大
    const Max = 1000
    const NumReceivers = 10
    wg := sync.WaitGroup{}
    wg.Add(NumReceivers)

    dataCh := make(chan int, 10)

    // the sender
    go func() {
        for {
            if value := rand.Intn(Max); value == 0 {
                close(dataCh)
                return
            } else {
                dataCh <- value
            }
        }
    }()
    // receivers
    for i := 0; i < NumReceivers; i++ {
        go func(i int) {
            defer wg.Done()
            for value := range dataCh {
                log.Println(value)
            }
            fmt.Printf("i = %v 结束\r\n", i)
        }(i)
    }
    wg.Wait()
}

//一个接收者，多个发送者，唯一的接收者通过关闭额外的channel通道表示‘请停止发送值到channel’
//这是一个比上面较复杂的情况。我们不能为阻止数据传输让接收者关闭数据channel，
//这样违反了channel关闭的原则，但是我们可以通过关闭额外的信号channel去通知发送者停止发送值。

func main5() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    const Max = 10
    const NumSenders = 10

    wg := sync.WaitGroup{}
    wg.Add(1)

    //----
    dataCh := make(chan int)
    stopCh := make(chan struct{})

    //many senders
    for i := 0; i < NumSenders; i++ {
        go func() {
            for {
                select {
                case <-stopCh:
                    return
                case dataCh <- rand.Intn(Max):
                }
            }

        }()
    }

    //one receiver
    go func() {
        defer wg.Done()
        for value := range dataCh {
            if value == Max-1 {
                log.Println(value)
                log.Println("关闭")
                close(stopCh)
                return
            }
            log.Println(value)
        }
    }()

    wg.Wait()
}

//M个接收者，N个发送者，任何一个通过中间人关闭信号channel表示‘让我们结束游戏吧’
//这是最复杂的情况。我们不能让任何一个发送者和接收者关闭数据channel。
//我们也不能让任何一个接收者关闭信号channel通知所有的接收者和发送者结束游戏。
//其中任何一种方式都打破了关闭原则。然而，我们可以引入一个中间角色去关闭信号channel。
//在下面例子中有一个技巧是如何使用try-send操作去通知中间人关闭信号通道。

func main6() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    const Max = 10000
    //接收者
    const NumReceivers = 10
    //发送者
    const NumSenders = 10

    wg := sync.WaitGroup{}
    wg.Add(NumReceivers)

    //数据
    dataCh := make(chan int)
    //关闭chan
    stopCh := make(chan struct{})
    //中间信号
    toStop := make(chan string, 1)
    var stopStr string

    go func() {
        time.Sleep(5 * time.Second)
        //中间信息
        stopStr = <-toStop
        close(stopCh)
    }()

    //发送者
    for i := 0; i < NumSenders; i++ {
        go func(id string) {
            for {
                value := rand.Intn(Max)
                //不发送了
                if value == 0 {
                    select {
                    case toStop <- "sender-" + id:
                    default:
                    }

                    return
                }

                select {
                case <-stopCh:
                    return
                //发送
                case dataCh <- value:

                }
            }
        }(strconv.Itoa(i))
    }

    //接收者
    for i := 0; i < NumReceivers; i++ {
        go func(id string) {
            defer wg.Done()
            for {
                select {
                //其中一个关闭也退出
                case <-stopCh:
                    return
                case data := <-dataCh:
                    if data == Max-1 {
                        select {
                        case toStop <- "sender-" + id:
                        default:
                        }
                        return
                    }
                    log.Println(data)

                }
            }
        }(strconv.Itoa(i))
    }

    wg.Wait()
    log.Println("stop by", stopStr)
}

//我们也可以设置toStop的缓冲大小是发送者和接收者之和。那样我们就不需要try-send的select块去通知中间人。?
func main() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    const Max = 100
    //接收者
    const NumReceivers = 10
    //发送者
    const NumSenders = 10

    wg := sync.WaitGroup{}
    wg.Add(NumReceivers)

    //数据
    dataCh := make(chan int)
    //关闭chan
    stopCh := make(chan struct{})
    //中间信号
    toStop := make(chan string, NumSenders+NumReceivers)
    var stopStr string

    go func() {
       /* for {
            time.Sleep(time.Second * 2)
            if len(toStop) == NumSenders+NumReceivers {
                fmt.Println("fuck")
                fmt.Println(NumSenders+NumReceivers)
                //中间信息*/
                stopStr = <-toStop
                close(stopCh)
              /*  return
            }
        }*/
    }()

    //发送者
    for i := 0; i < NumSenders; i++ {
        go func(id string) {
            for {
                value := rand.Intn(Max)
                //不发送了
                if value == 0 {
                    toStop <- "sender-" + id
                    return
                }

                select {
                case <-stopCh:
                    return
                //发送
                case dataCh <- value:

                }
            }
        }(strconv.Itoa(i))
    }

    //接收者
    for i := 0; i < NumReceivers; i++ {
        go func(id string) {
            defer wg.Done()
            for {
                select {
                //其中一个关闭也退出
                case <-stopCh:
                    return
                case data := <-dataCh:
                    if data == Max-1 {
                        toStop <- "sender-" + id
                        return
                    }
                    log.Println(data)

                }
            }
        }(strconv.Itoa(i))
    }

    wg.Wait()
    log.Println("stop by", stopStr)
}
