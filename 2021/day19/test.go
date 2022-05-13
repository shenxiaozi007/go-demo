package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func main() {
	fmt.Println()
}
