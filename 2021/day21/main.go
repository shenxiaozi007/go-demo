package main

import (
	"fmt"
	"time"
)

func main()  {
	stop := make(chan struct{})
	go func() {
		fmt.Println("test1")
		for  {
			select {
				case <-stop:
					fmt.Println("测试1")
					return
			}
		}


	}()

	fmt.Println("主线程退出")
	for  {
		time.Sleep(5 * time.Second)
		<-stop
	}
}
