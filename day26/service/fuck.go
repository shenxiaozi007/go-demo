package main

import "fmt"

type tt struct {
}

func main1() {
	t()
}

func t() {
	a := new(tt)

	b := new(tt)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a == b)

}