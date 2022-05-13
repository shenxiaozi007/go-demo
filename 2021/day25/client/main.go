package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

//主函数
func main() {
	//1 连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	//调用方法
	ret := 0
	err2 := conn.Call("Rect.Area", Params{20, 100}, &ret)

	if err2 != nil {
		log.Fatal(err2)

	}

	fmt.Println("面积", ret)
	//周长
	err3 := conn.Call("Rect.Perimeter", Params{50, 60}, &ret)

	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
}