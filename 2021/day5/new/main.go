package main

import "fmt"

//int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string
func main() {

	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)
}
