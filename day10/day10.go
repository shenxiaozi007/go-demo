package main

import (
    "context"
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

//https://www.zhihu.com/question/450188866
func main4() {
    c := make(chan int, 10)

    wg := sync.WaitGroup{}
    ctx, cancel := context.WithCancel(context.Background())
    //只有一个关闭
    go func() {
        //先停止写
        time.Sleep(time.Second * 2)
        cancel()
        //再关闭
        close(c)
    }()

    //多个写
    for i := 0; i <= 10; i++ {
        go func(ctx context.Context) {
            select {
            case <-ctx.Done():
                fmt.Println("主动关闭")
                return
            case c <- i:
            }
        }(ctx)
    }

    //多个读
    for i := 0; i <= 10; i++ {
        wg.Add(1)
        go func(i int) {
            //最后执行
            defer wg.Done()
            time.Sleep(time.Second * 3)
            for v := range c {
                fmt.Printf("输出中：%s \r\n", v)
            }
            fmt.Printf("输出完毕:%s \r\n", i)
        }(i)
    }

    wg.Wait()
    fmt.Println("执行完毕")
}

func main2() {
    //len表示实际的切片内元素数量；
    //cap表示切片的真实底层数量；
    test1 := make([]int, 299, 300)
    test2 := make([]int, 1199, 1200)
    //cap如果小于1024，那么就扩大为2len(这里是排除掉新元素的原len)
    //cap如果大于1024，那么久扩大为1.25len(这里是排除掉新元素的原len)
    test1 = append(test1, 11)
    test2 = append(test2, 11)
    fmt.Println(test1,len(test1),cap(test1))
    fmt.Println(test2,len(test2), cap(test2))
}
var x int32
var wg sync.WaitGroup
func main3()  {

    wg.Add(2)
    go add()
    go add()

    wg.Wait()

    fmt.Println(x)
}

func add()  {
    for i := 0; i<5000; i++ {
        //原子操作
        atomic.AddInt32(&x, 1)
    }

    defer wg.Done()
}

func main1() {
    fmt.Println(getNumber())
}

func getNumber() int {
    var i int

    go func() {
        i = 5
    }()
    return i
}

//关于chan缓存满时的读取顺序
func main() {

    cn := make(chan int, 5)
    
    go func() {
        for i := 0; i < 15; i++ {
            cn <- i
        }
        //写完关闭
        close(cn)
    }()

    wg := sync.WaitGroup{}
    wg.Add(2)
    lock := sync.Mutex{}
    for i := 0; i < 2; i++ {
        go func() {
            lock.Lock()
            defer lock.Unlock()
            defer wg.Done()
            for  {
                select {
                case i, ok := <-cn:
                    if !ok {
                        return
                    }
                    fmt.Println(i)

                }
            }
        }()
    }


    wg.Wait()
    fmt.Println("执行完毕")
}

