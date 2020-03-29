package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	var b = []int{1, 2, 3}
	fmt.Println(a, b)
	fmt.Printf("a:%T b:%T", a, b)

}
