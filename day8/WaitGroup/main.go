package main

import (
	"sync"
	"fmt"
	"runtime"
	"time"
)

//有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
// 类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

var x int64
var wg sync.WaitGroup

//互斥锁
//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
var lock sync.Mutex

//读写互斥锁
//互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。
//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 100000; i++ {
		x = x + 1
	}
	wg.Done()
}


//互斥锁
func mutexAdd() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

//读写互斥锁
func write() {
	lock.Lock()   // 加互斥锁
	//rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	//rwlock.Unlock()                   // 解写锁
	lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	lock.Lock()                  // 加互斥锁
	//rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	//rwlock.RUnlock()             // 解读锁
	lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go mutexAdd()
	go mutexAdd()
	wg.Wait()
	fmt.Println(x)

	//互斥锁

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

