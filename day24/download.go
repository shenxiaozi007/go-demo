package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//下载文件
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	fmt.Println(resp)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:/img/" + filename

	//写入数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	fmt.Println(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}

//定义变量
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	//监控协程
	chanTask chan string
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	test := DownloadFile("https://uploadfile.bizhizu.cn/up/9b/39/e5/9b39e505ca198208c1132c9477fd6844.jpg.230.350.jpg", "1.jpg")
	fmt.Println(test)

	//初始化
	chanImageUrls := make(chan string, 1000)
	chanTask := make(chan string, 26)
	//爬虫携程


}

func GetPagestr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
}
