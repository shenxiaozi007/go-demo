package main

import "fmt"

func main() {
	s1 := "hello"
	byteArray := []byte(s1)
	fmt.Println(byteArray)

	s2 := ""

	for i := len(byteArray) - 1; i >= 0; i-- {
		s2 += string(byteArray[i])
	}
	fmt.Println(s2)
}
