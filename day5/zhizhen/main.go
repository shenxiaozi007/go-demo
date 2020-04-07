package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Printf("a: %d ptr:%p \n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	c := &b
	fmt.Println(c)
	fmt.Println(*c)

	o := 10
	x := modify1(o)
	fmt.Println(o)
	y := modify2(&o)
	fmt.Println(x, y)
	fmt.Println(o)

}

func modify1(x int) int {
	x = 100
	return x
}

func modify2(x *int) int {
	*x = 100
	return *x
}
