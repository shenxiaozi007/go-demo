package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

type T struct {
	n int
}

func main1() {
	m := make(map[int]T)
	tmp := m[0]
	tmp.n = 1
	fmt.Println(m[0].n)
}

type X struct{}

func (x *X) test() {
	println(x)
}

func main2() {
	var a *X
	a.test()
	var b = X{}
	// X{}.test()
	b.test()
}

//X{} 是不可寻址的，不能直接调用方法。知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。

var count int

func Count(lock *sync.Mutex) {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}

func putStr(intTmp chan int, w *sync.WaitGroup) {
	for ch := range intTmp {
		fmt.Println(ch)
		if ch >= 9 {
			break
		}
	}
	w.Done()
}

func main3() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		c := count
		if c >= 10 {
			break
		}

	}

	fmt.Println("fuck")
	int_chan_list := make(chan int, 10)
	for m := 0; m < 10; m++ {
		int_chan_list <- m
	}

	w := sync.WaitGroup{}
	w.Add(1)
	go putStr(int_chan_list, &w)
	w.Wait()
}

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)

	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	fmt.Println("Hello World")
}


