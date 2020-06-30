package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//每日练习 数组
	var a [3]int
	a[1] = 1
	a[2] = 2
	var b [3]string
	b[1] = "3232"
	fmt.Println(a)
	fmt.Printf("%T", a)
	//循环
	for i := 0 ; i < 3 ; i ++ {
		fmt.Println(a[i])
	}

	for val, key := range a {
		fmt.Println(val,key)
	}
	//
	for v, k := range b {
		fmt.Println(v,k)
	}
	//每日练习 切片 切片是引用类型
	//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。

	var a1 []int
	//a1 = make([]int, 4)
	a2 := []string{}
	a3 := make([]bool, 2)
	fmt.Printf("test--%#v---", a1)
	fmt.Printf("test--%#v---", a2)
	//a1[3] = 1
	a1 = append(a1, 1)
	a2 = append(a2,"test")
	a3 = append(a3, true)
	fmt.Println(a1,a2,a3)

	a4 := a3
	a4[0] = true
	fmt.Println(a1,a2,a3,a4)

	//每日练习 map[KeyType]ValueType
	//map是引用类型 需要make初始化后才可以使用

	var map1 map[int]string
	fmt.Printf("%#v---\n", map1)
	map1 = make(map[int]string, 10)
	map1[1] = "3232"

	fmt.Printf("%#v---\n", map1)
	//

	map2 := make(map[string][]string)

	mapSlice := []string{"1", "2"}
	map2["test"] = mapSlice

	fmt.Printf("%#v---\n", map2)

	fmt.Printf("---%#v---\n", map2["test2"])
	//判断map2是否存在
	val, ok := map2["test2"]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("不存在")
	}

	//按顺序循环读取
	mapTest1 := make(map[string]int, 100)

	rand.Seed(time.Now().UnixNano())
	sliceList := []string{}

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("test%02d", i)
		value := rand.Intn(100)
		mapTest1[key] = value
	}
	fmt.Println(mapTest1)
	//赋值到slice
	for key, _ := range mapTest1 {
		sliceList = append(sliceList, key)
	}
	fmt.Println(sliceList)
	//排序
	sort.Strings(sliceList)
	fmt.Println(sliceList)
	//按排序获取map值
	for _, val := range sliceList {
		mapStr := mapTest1[val]
		fmt.Printf("%v---%v---\n",val, mapStr)
	}

}
