package models

import (
	"github.com/astaxie/goredis"
)

var (
	client goredis.Client
)

const (
	//ReadUrls  存放已经爬过的url
	ReadUrls = "Read_Urls"
	// BookUrls 存放每本书的url
	BookUrls = "book_urls"

	//PAGE_URLS 生成每一页的url
	PAGE_URLS = "page_urls"
)

func ConnectRedis(addr string) {
	client.Addr = addr
}

//PutInQue 从页面中解析出来的book的url 放入队列中
func PutInQue(que, url string) {
	//fmt.Println(que)
	err := client.Lpush(que, []byte(url))
	if err != nil {
		panic(err)
	}
}

// GetUrl 从队列中取出一个URL
func GetUrl(que string) string {
	//fmt.Println("111111")
	b, err := client.Lpop(que)
	if err != nil {
		panic(err)
	}
	url := string(b)
	return url
}

// IsInReadQue 判断url 是否已经爬取过
func IsInReadQue(url, que string) bool {
	//fmt.Println("222222")
	b, err := client.Hexists(que, url)
	if err != nil {
		return false
	}
	return b
}
