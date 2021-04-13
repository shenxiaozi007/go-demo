package main

import (
	"fmt"
	"sync"
)

//闭包
func a() func() int {
	i := 0

	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func add(add int) func(int) int {
	return func(i int) int {
		add += i
		return add
	}
}

/*func b() func() int {
	i := 0

}*/

func main() {
	//tmp1 := add(10)
	//fmt.Println(tmp1(1), tmp1(2))
	//
	//tmp2 := add(10)
	//fmt.Println(tmp2(1), tmp2(2))
	//c()
	//c()
	//c()
	wg := sync.WaitGroup{}
	wg.Add(2)
	output := make(chan int)
	output2 := make(chan int)

	go func(output chan int) {
		for i := 0; i < 5; i++ {
			output <- i + 100
		}
		close(output)
	}(output)

	go func(output2 chan int) {
		for i := 0; i < 5; i++ {
			output2 <- i
		}
		close(output2)
	}(output2)

	go func() {
		for {
			select {
			case value, ok := <-output:
				if !ok {
					fmt.Println("fuck11")
					wg.Done()
				}
				fmt.Println(value)
				break
			case value1, ok := <-output2:
				if !ok {
					fmt.Println("fuck22")
					wg.Done()
					break
				}
				fmt.Println(value1)
				break
			}
		}
	}()
	wg.Wait()

}
