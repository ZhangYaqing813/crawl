package models

import (
	"github.com/astaxie/goredis"
)

var (
	client goredis.Client
)

const (
	PAGE_URLS = "page_urls"
	BOOK_URLS = "book_urls"


)

func ConnectRedis(addr string)  {
	client.Addr = addr
}

func PutInQue(url string)  {
	client.Lpush()
}

