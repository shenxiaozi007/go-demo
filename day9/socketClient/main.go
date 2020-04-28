package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

//建立与服务端的链接
//进行数据收发
//关闭链接

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("err :", err)
		return
	}

	defer conn.Close()

	//收
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入

		inputInfo := strings.Trim(input, "\r\n")

		//如果为Q就退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		_, err := conn.Write([]byte(inputInfo))

		if err != nil {
			fmt.Println("write failed, err", err)
			return
		}

		buf := [512]byte{}

		//读取服务返回的
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("rece failed , err", err)
			return
		}

		fmt.Println(string(buf[:n]))

	}



}