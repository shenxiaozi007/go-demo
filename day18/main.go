package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//一
func main1() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()

		}()
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

//二
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

//三
func main3() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("a") }()
	defer func() { fmt.Println("b") }()
	defer func() { fmt.Println("c") }()
	panic("触发异常")
}

func main4()  {
	fmt.Println(test(2))
}

func test(i int) (str string) {
	str = "11"
	if i == 1 {
		return "22"
	}
	return
}
//四

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//五
func main5() {

	ua := &UserAges{}
	ua.ages = make(map[string]int)

	for i := 0; i < 10000; i++ {
		str1 := fmt.Sprintf("test+%v", i)
		go func() {
			ua.Add(str1, i)
		}()
	}

	for i := 0; i < 10000; i++ {
		str2 := fmt.Sprintf("test+%v", i)
		go func() {
			fmt.Println(ua.Get(str2))
		}()
	}
	time.Sleep(1 * time.Second)
}

//六
func main6() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

//七
func main7() {
	a := 1
	b := 2

	defer calc7("1", a, calc7("10", a, b))
	a = 0
	defer calc7("2", a, calc7("20", a, b))
	b = 1
}

func calc7(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//八
func main8() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//[0 0 0 0 0 1 2 3]
}


//九 下面的迭代会有什么问题？
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{},len(set.s))
	go func() {
		set.RLock()

		for elem,value := range set.s {
			ch <- elem
			println("Iter:",elem,value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main()  {

	th := threadSafeSet{
		s:[]interface{}{"1","2"},
	}
	fmt.Printf("%v", <-th.Itest())
}

func (s *threadSafeSet) Itest() <-chan interface{} {
	ch := make(chan interface{}, len(s.s))
	//ch := make(chan interface{})
	go func() {
		s.RLock()

		for elem, value := range s.s {
			ch <- elem
			fmt.Println("Iter1:", elem, value)
		}

		close(ch)
		s.RUnlock()
	}()

	return ch
}

