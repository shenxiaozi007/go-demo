package main

import "fmt"


//闭包
func a() func() int {
	i := 0

	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func add(add int) func(int) int {
	return func(i int) int {
		add += i
		return add
	}
}
/*func b() func() int {
	i := 0

}*/

func main()  {
	f := a()
	f()
	f()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := add(10)
	fmt.Println(tmp2(1), tmp2(2))
	//c()
	//c()
	//c()

	a()
	//c2()
	//c2()

	a()
	a()

}

