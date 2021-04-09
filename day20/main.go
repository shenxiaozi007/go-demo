package main

import (
	"fmt"
	"sync"
	"time"
)

//写入
func put(num chan int, wg *sync.WaitGroup, lock *sync.Mutex) {
	//defer wg.Done()
	for i := 0; i < 10; i++ {
		go func(i int) {
			//lock.Lock()
			//defer lock.Unlock()
			defer wg.Done()
			fmt.Println(i)
			num <- i
		}(i)
	}
}


func main1() {
	//runtime.GOMAXPROCS(1)
	lock := new(sync.Mutex)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	num := make(chan int)

	go put(num, wg, lock)
	//go out(num)

	go func(lock *sync.Mutex) {
		for _ = range num {
			//lock.Lock()
			//defer lock.Unlock()
			fmt.Println(11)
		}
	}(lock)
	wg.Wait()
	close(num)
	time.Sleep(2 * time.Second)

}

func main2() {
	//lock := sync.Mutex{}

	//无缓冲
	ch := make(chan int)
	//单
	go func() {
		for i := 0; i < 10; i ++{
			//加锁
			//lock.Lock()
			//defer lock.Unlock()
			ch <- i
		}
		close(ch)
	}()
	//for i:=0; i < 10; i++ {
	//	num := <- ch
	//	fmt.Println(num)
	//}
	for num := range ch{
		fmt.Println(num)
	}
}

func main()  {
	//wg := sync.WaitGroup{}
	ch := make(chan int, 20)
	num := 0
	for i := 0; i < 20; i++{
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			num += 1
			//time.Sleep(2 * time.Second)
			//fmt.Println(num, i)
			ch<- num
			if num == 20 {
				//关闭通道
				close(ch)
			}
		}()
	}

	for value := range ch {
		fmt.Println(value)
	}
	//wg.Wait()
}

