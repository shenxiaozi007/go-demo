package main

import (
	"fmt"
	"runtime"
	"sync"
)

//面试题一
func main1() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()

		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("B: ", i)
			wg.Done()
		}()

	}

	wg.Wait()
}

//面试题二
type People struct {

}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show B")
}

type Teacher struct {
	People
}

func (t *Teacher)  ShowB() {
	fmt.Println("teacher show B")
}

func main2() {
	t := Teacher{}
	t.ShowA()
}

//面试题三
func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("a") }()
	defer func() { fmt.Println("b") }()
	defer func() { fmt.Println("c") }()
	panic("触发异常")
}