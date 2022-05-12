package main

func test1(i *int) {

}

func main() {
	i := 5

	t := test1

	t(&i)

}