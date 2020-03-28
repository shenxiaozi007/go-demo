package main

import "fmt"

func main() {
	ifDemo()
	forDemo2()
}

func ifDemo() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo2() {
	i := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//while
	for i < 10 {
		fmt.Println("while")
		i++
	}

	//无限循环
	for {
		if i > 20 {
			break
		}
		fmt.Println("for")
		i++
	}
	//switch
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4:
		fmt.Println("偶数")
	}
	//
	age := 30

	switch {
	case age < 25:
		fmt.Println("小于25")
	case age > 25:
		fmt.Println("大于25")
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			goto breakTag
		}
		fmt.Println(i)
	}

	//标签
breakTag:
	fmt.Println("tag2")

}
