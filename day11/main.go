package main

import (
	"fmt"
	"github.com/huangxinchun/go-demo/testpackage"
	"github.com/huangxinchun/go-demo"
)

func main() {
	go_demo.New()
	testpackage.TestNew()
	fmt.Println("main")
}
