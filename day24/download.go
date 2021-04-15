package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//下载文件
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "D:/img/" + filename

	//写入数据
	err = ioutil.WriteFile(filename, bytes, 0666)
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
	//test := DownloadFile("https://uploadfile.bizhizu.cn/up/9b/39/e5/9b39e505ca198208c1132c9477fd6844.jpg.230.350.jpg", "1.jpg")
	//fmt.Println(test)

	//1.初始化管道
	chanImageUrls = make(chan string, 1000)
	chanTask = make(chan string, 26)
	//2.爬虫携程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrl("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	//3, 任务统计协程，统计26个任务是否都完成。完成就关闭管道
	waitGroup.Add(1)
	go CheckOk()

	//4. 下载协程， 从管道读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

func DownloadImg() {
	for url := range chanImageUrls{
		filename := GetFilenameFromurl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

//截取url名字
func GetFilenameFromurl(url string) (filename string) {
	//返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	//切出来
	filename = url[lastIndex+1:]
	//时间错解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))

	filename = timePrefix + "_" + filename
	return
}
func CheckOk() {
	var count int
	for  {
		url := <-chanTask
		fmt.Printf("%s 完成了爬虫任务 \n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
}

//遍历链接到管道里面
func getImgUrl(url string)  {
	urls := getImgs(url)
	fmt.Println(urls)
	for _, urlVaule := range urls {
		chanImageUrls <- urlVaule
	}
	//fmt.Println(chanImageUrls)
	//标记协程完成
	//完成一个写一个数据
	chanTask <- url
	waitGroup.Done()
}

//获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPagestr(url)

	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)

	fmt.Printf("共找到%d条结果 \n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

func GetPagestr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	//读取页面信息
	pagesBytes, err := ioutil.ReadAll(resp.Body)

	HandleError(err, "ioutil.ReadAll")
	//字节转字符串
	pageStr = string(pagesBytes)

	return pageStr
}
