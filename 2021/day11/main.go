package main

import (
	"fmt"

	go_demo "github.com/huangxinchun/go-demo"
	"github.com/huangxinchun/go-demo/testpackage"
)

func main() {
	go_demo.New()
	testpackage.TestNew()
	fmt.Println("main")
}
