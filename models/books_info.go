package models

import (
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/httplib"
	"regexp"
	"strconv"
)

type BooksInfo struct {
	Id		int
	Name 	string
	Author  string
	Score	float32
	Isbn	int
}

// GetTotalPages 获取当前页面中最大的页面数量
func GetTotalPages(bookHtml string) int {
	// expression 正则表达式内容
	expression := `<a href="/tag/编程.*?type=T" >(\d+?)</a>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(bookHtml,-1)
	if len(result) == 0 {
		return 0
	}
	pagesNumber,err :=strconv.Atoi( result[len(result)-1][1])
	if err != nil {
		panic(err)
	}
	return pagesNumber
}

//GetBookUrl 获取页面中book 的URL信息。
func GetBookUrl(bookHtml string) []string {
	// expression 正则表达式内容 提取book 的URL
	expression := `<a href="(https://book.douban.com/.*?/)" `
	var urls []string
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(bookHtml,-1)
	for i :=0; i<len(result);i++ {
		urls = append(urls,result[i][1])
	}
	return urls
}

