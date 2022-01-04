package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("監控推出")
				return
			default:
				fmt.Println("監控中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了。统计监听停止")

	stop <- true

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//多个协程退出
func main2() {
	stopSingal := make(chan bool)

	//循环
	for i := 1; i < 5; i++ {
		go func(ch chan bool, number int) {
			for {
				select {
				case v := <-ch:
					// 仅当 ch 通道被 close，或者有数据发过来(无论是true还是false)才会走到这个分支
					fmt.Printf("监控器%v，接收到通道为：%v 监控结束。\n", number, v)
					return
				default:
					fmt.Printf("监控器%v，正在监控中...\n", number)
					time.Sleep(2 * time.Second)
				}
			}
		}(stopSingal, i)
	}
	time.Sleep(4 * time.Second)

	stopSingal <- true
	stopSingal <- true

	time.Sleep(10 * time.Second)

	//关闭所有的goroutine

	close(stopSingal)
	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出！！")
}

func main3() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i < 5; i++ {
		go func(ctx context.Context, number int) {
			for {
				select {
				case v := <-ctx.Done():
					fmt.Printf("监控器 %v，接受到通道值为：%v, 监控结束。\n", number, v)
					return
				default:
					fmt.Printf("监控器%v，正在监控中...\n", number)
					time.Sleep(2 * time.Second)
				}
			}
		}(ctx, i)
	}

	time.Sleep(1 * time.Second)

	//关闭所有 goroutine
	cancel()

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)

	fmt.Println("退出主程序")
}

func main4() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt.....")
				return

			default:
				fmt.Printf("send message : %d \n", <-messages)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}

func main() {

	message := make(chan int, 10)

	for i := 0; i < 10; i++ {
		message <- i
	}

	ctx5, _ := context.WithTimeout(context.Background(), 4*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	// consumer

	go main5test(ctx5, message)

	time.Sleep(10 * time.Second)
	close(message)
	cancel()
	// time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}

}

func main5test(ctx context.Context, message chan int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("child process interrupt...")
			return
		default:
			fmt.Printf("send message: %d\n", <-message)
		}
	}
}
