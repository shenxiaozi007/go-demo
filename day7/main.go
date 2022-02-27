package main

import "runtime/debug"

func main1() {
    var a [1]int
    c := a[:]
    println(c)
}

func main2()  {
    slice := make([]string, 2, 4)
    example(slice, "hello", 10)
}

func example(slice []string, str string, i int) {
    debug.PrintStack()
}

func main() {

}