package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {}

// Rpc 服务端 求矩形面积

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

//周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

//主函数
func main() {
	//注册服务

	rect := new(Rect)

	//1注册一个rect服务
	rpc.Register(rect)

	//2 服务处理绑定到http协议上
	rpc.HandleHTTP()
	//监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("tests")
}
