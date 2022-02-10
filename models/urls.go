package models

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "github.com/astaxie/beego/httplib"
	"regexp"
	"strconv"
)

func GetHtmlStr(url string) string {

	resp := httplib.Get(url)
	//得每页的html数据
	strHtml, err := resp.String()
	if err != nil {
		panic(err)
	}
	return strHtml
}

// GetTotalPages 获取当前页面中最大的页面数量
func GetTotalPages(bookHtml string) int {
	// expression 正则表达式内容
	expression := `<a href="/tag/编程.*?type=T" >(\d+?)</a>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(bookHtml, -1)
	//fmt.Println(result)
	if len(result) == 0 {
		return 0
	}
	pagesNumber, err := strconv.Atoi(result[len(result)-1][1])

	if err != nil {
		panic(err)
	}
	//fmt.Println("-------",pagesNumber)
	return pagesNumber
}

//GetBookUrl 获取页面中book 的URL信息。
func GetBookUrl(bookHtml string) []string {
	//fmt.Println(bookHtml)
	exp := `<p class='pl2'>没有找到符合条件的图书</p>`
	r := regexp.MustCompile(exp)
	result := r.FindAllStringSubmatch(bookHtml, -1)
	var urls []string
	if len(result) == 0 {
		// expression 正则表达式内容 提取book 的URL
		expression := `<a class="nbg" href="(https://book.douban.com/.*?/)"`

		reg := regexp.MustCompile(expression)
		result = reg.FindAllStringSubmatch(bookHtml, -1)
		for i := 0; i < len(result); i++ {
			urls = append(urls, result[i][1])
		}
	} else {
		return urls
	}
	//fmt.Println("GetBookUrl",urls)
	return urls
}
