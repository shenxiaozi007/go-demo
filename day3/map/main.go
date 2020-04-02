package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

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

	//刪除键值对
	delete(userInfo, "test1")
	fmt.Println(userInfo)

	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	randMap := make(map[string]int, 200)

	for i := 0; i < 10; i++ {
		keyMap := fmt.Sprintf("key %d", i)
		value := rand.Intn(100) //生成0~99的随机整数
		randMap[keyMap] = value
	}
	fmt.Println(randMap)
	//取出map中所有key存入切片
	var keys = make([]string, 0, 200)
	for key := range randMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)

	//按照排序后的key遍历
	for _, keyS := range keys {
		fmt.Println(keyS, randMap[keyS])
	}

	sliceMap := make(map[string][]string, 3)

	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}

	value = append(value, "北京", "上海")

	sliceMap[key] = value
	fmt.Println(sliceMap)
	
}
