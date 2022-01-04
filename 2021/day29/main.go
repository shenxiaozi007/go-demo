package main

func add(x, y int) (z int) {
	defer func() {
		println(z)
	}()

	z = x + y
	return z + 200
}


func addTwo(x, y int) (z int) {
	defer func() {
		z += 100
	}()

	z = x + y
	return
}


func main() {
	//println(add(1, 2)) //
	//println(addTwo(1, 2)) //
}

