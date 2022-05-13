package main

import "fmt"

//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["test"] = 100
	fmt.Println(b)

	c := make(map[string]int, 20)
	c["test"] = 111
	fmt.Println(c)
}
