package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)

	scoreMap["fuck"] = 1
	scoreMap["test"] = 2

	fmt.Println(scoreMap)

	//map也支持在声明的时候填充元素
	userInfo := map[string]string{
		"test1": "test2",
		"test3": "test4",
	}

	fmt.Println(userInfo)

	//判断某个键是否存在 Go语言中有个判断map中键是否存在的特殊写法，格式如下:
	v, ok := userInfo["test1"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}
	//Go语言中使用for range遍历map。
	for k, v := range userInfo {
		fmt.Println(k, v)
	}

}
