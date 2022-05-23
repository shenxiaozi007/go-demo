package main

import (
    "fmt"
    "math"
    "runtime"
    "sync"
    "sync/atomic"
)

type StrTest struct {
}

func NewStrTest() *StrTest {
    return new(StrTest)
}
func (st *StrTest) getStr(hayStack string, needle string) (data int) {
    hayByte := []byte(hayStack)
    needByte := []byte(needle)
    needCount := len(needByte)
    for k, v := range hayByte {
        nextK := 0
        nowK := 0
        countK := 0
        for _, n := range needByte {
            //判断是否相等
            if v == n {
                if needCount == 1 {
                    data = k
                    return
                } else if needCount > 1 {
                    countK++
                    fmt.Println(k, countK, needCount, string(v), string(n))
                    if countK == needCount {
                        data = k
                        return
                    }
                    if nextK != 0 {
                        //保证顺序
                        if nowK+1 != nextK {
                            data = -1
                        }
                    }
                    //判断下一个
                    nowK = k
                    nextK = nowK + 1
                    continue
                }

            }
        }
    }
    return
}
func main1() {
    fmt.Println("my fuck")

    testNewStr := NewStrTest()
    hayStack := "he333llo"
    needle := "l3l"
    str := testNewStr.getStr(hayStack, needle)
    fmt.Println(str)
}

func main2() {
    var a uint = 1
    var b uint = 2
    test := math.Pow(2, 64)

    fmt.Println(test)
    fmt.Println(a - b)
}

func Cat(wg *sync.WaitGroup, catCount uint64, catch chan struct{}, dog chan struct{}) {
    defer wg.Done()
    for {
        if catCount >= 100 {
            return
        }
        select {
        case <-catch:
            fmt.Println("cat")
            dog <- struct{}{}
            //加1
            atomic.AddUint64(&catCount, 1)
        }
    }
    fmt.Println(catCount)
}

func Dog(wg *sync.WaitGroup, dogCount uint64, dogch chan struct{}, fishCh chan struct{}) {
    defer wg.Done()
    for {
        if dogCount >= 100 {
            return
        }
        select {
        case <-dogch:
            fmt.Println("dog")
            fishCh <- struct{}{}
            //加1
            atomic.AddUint64(&dogCount, 1)
        }
    }
    fmt.Println(dogCount)
}

func Fish(wg *sync.WaitGroup, fishCount uint64, fishch chan struct{}, catCh chan struct{}) {
    defer wg.Done()

    for {
        if fishCount >= 100 {
            return
        }
        select {
        case <-fishch:
            fmt.Println("fish")
            catCh <- struct{}{}
            //加1
            atomic.AddUint64(&fishCount, 1)
        }
    }
    fmt.Println(fishCount)
}

func main5() {
    var catCount uint64
    var dogCount uint64
    var fishCount uint64
    catCh := make(chan struct{}, 1)
    dogCh := make(chan struct{}, 1)
    fishCh := make(chan struct{}, 1)
    wg := &sync.WaitGroup{}
    wg.Add(3)
    go Cat(wg, catCount, catCh, dogCh)
    go Dog(wg, dogCount, dogCh, fishCh)
    go Fish(wg, fishCount, fishCh, catCh)
    //第一个是cat
    catCh <- struct{}{}
    wg.Wait()
}

func main6() {
    arr := make([]int, 0)
    for i := 0; i < 2000; i++ {
        fmt.Println("len，为", len(arr), "cap 为", cap(arr))
        arr = append(arr, i)
    }

    runtime.GC()
}

func main() {
    test := 1 << 3
    fmt.Println(test)
}
