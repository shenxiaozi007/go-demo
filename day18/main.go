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

func main9()  {

	th := threadSafeSet{
		s:[]interface{}{"1","2"},
	}
	fmt.Printf("%v", <-th.Itest())
}

//十

type People1 interface {
	Speak(string) string
}

type Student1 struct {}

func (stu *Student1) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "you are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestPeople(p interface{})  {
	switch p.(type) {
	case People1:
		fmt.Println("实现了people1")
	case Student1:
		fmt.Println("实现了Student1")
	case notifier:
		fmt.Println("notifier")
	case User:
		fmt.Println("User")
	}
}
func main10() {
	peo := Student1{}
	TestPeople(peo)
	think := "bitch"
	fmt.Println(peo.Speak(think))
}


// 定义一个接口

type notifier interface {
	Notify()
}

type User struct {
	Name string
	Age  int
}

func (s *User) Notify1() {
	fmt.Println("name is: ", s.Name)
}

func (s User) Notify() {
	fmt.Println("name is111: ", s.Name)
}

func notifination(n notifier) {
	n.Notify()
}

func notifination1(n User) {
	n.Notify1()
	n.Notify()
}

func main11() {
	u := User{"james", 33}
	TestPeople(u)
	//notifination(u)
	notifination1(u)
}

//12

func main12() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

//13
func main13() {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
}

//14
func main14()  {
	s1 := []int{1, 2, 3}
	s2 := []int{2, 3}
	s2 = append(s1, s2...)
	fmt.Println(s2)
}

//15
func Foo(x interface{}) {
	fmt.Printf("%#v", x)
	fmt.Println(x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main15() {
	var x *int = nil
	Foo(x)
}

//16
//考点：函数返回值类型
//nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。报:cannot use nil as type string in return argument.
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}

func main16()  {
	intMap := map[int]string{
		1 : "a",
		2 : "bb",
	}

	v, err := GetValue(intMap, 2)
	fmt.Println(v, err)
}


//17
func test1(x int) (func(),func())  {
	return func() {
			println(x)
			x+=10
		}, func() {
			println(x)
		}
}

func main()  {
	a, b := test1(100)
	a()
	b()
}


