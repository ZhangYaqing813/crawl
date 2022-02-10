package models

import (
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/httplib"
	"regexp"
	"strconv"
)

type BooksInfo struct {
	Id     int
	Name   string
	Author string
	Score  float64
	Isbn   int
}

func GetBookName(strHtml string) (name string) {
	expression := `<span property="v:itemreviewed">(.*?)</span>`
	expression2 := `<span class="pl">副标题:</span> 从入门到实践<br/>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(strHtml, -1)

	reg1 := regexp.MustCompile(expression2)
	result1 := reg1.FindAllStringSubmatch(strHtml, -1)

	if len(result) == 0 {
		return ""
	}
	if len(result1) == 0 {
		name = result[0][1]
	} else {
		name = result[0][1] + ":" + result1[0][1]
	}
	return name
}

func GetBookAuthor(strHtml string) (Author string) {
	expression := `<a class="" href="/search/.*?">(.*?)</a>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(strHtml, -1)
	//fmt.Println(result)
	if len(result) == 0 {
		return ""
	}
	return result[0][1]
}

func GetBookScore(strHtml string) (score float64) {
	expression := `<strong class=".*?"v:average">(.*?)</strong>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(strHtml, -1)
	//fmt.Println(result)
	if len(result) == 0 {
		return 0
	}
	score, _ = strconv.ParseFloat(result[0][1], 64)
	return score
}

func GetBookIsbn(strHtml string) (isbn int) {
	expression := `<span class="pl">ISBN:</span>(\d+?)<br/>`
	reg := regexp.MustCompile(expression)
	result := reg.FindAllStringSubmatch(strHtml, -1)
	//fmt.Println(result)
	if len(result) == 0 {
		return 0
	}
	isbn, _ = strconv.Atoi(result[0][1])
	return isbn
}
