package main

import (
	"net"
	"fmt"
	"bufio"
)


//1.监听端口
//2.接收客户端请求建立链接
//3.创建goroutine处理链接。


//处理链接
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte
		fmt.Println("等待发送", conn.LocalAddr(), conn.RemoteAddr())
		n, err := reader.Read(buf[:]) //读取数据

		if err != nil{
			fmt.Println("read from client failed, err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到的数据是", recvStr)

		conn.Write([]byte(recvStr))//发送数据
	}
}

func main()  {

	listen, err := net.Listen("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("linstn failed err", err)
		return
	}

	for {
		conn, err := listen.Accept() //建立連接

		if err != nil {
			fmt.Println("accept failed err:", err)
			continue
		}
		go process(conn)

	}
}