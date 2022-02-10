package controllers

import (
	"crawl_books/models"
	"fmt"
	_ "fmt"
	"github.com/astaxie/beego"
	"strconv"
	"sync"
	"time"
)

type BooksController struct {
	beego.Controller
}

func init() {
	models.ConnectRedis("127.0.0.1:6379")
}

func (c *BooksController) BooksInfo() {
	wg := sync.WaitGroup{}
	// startPage 用于控制页数
	startPage := ""
	startUrl := "https://book.douban.com/tag/%E7%BC%96%E7%A8%8B?start=20&type=T"

	// 获取该分类下一共有多少页
	pages := models.GetTotalPages(models.GetHtmlStr(startUrl))
	for i := 0; i < pages; i++ {
		startPage = strconv.Itoa(i * 20)
		pagesUrl := "https://book.douban.com/tag/%E7%BC%96%E7%A8%8B?start=" + startPage + "&type=T"
		//fmt.Println(pagesUrl)
		models.PutInQue(models.PAGE_URLS, pagesUrl)
	}
	var book models.BooksInfo

	// 获取每一个页面中所有book 的url 并存入BookUrls 队列中
	for {
		pageUrl := models.GetUrl(models.PAGE_URLS)
		if len(pageUrl) == 0 {
			break
		}
		//fmt.Println("ccccccc",pageUrl)
		if !models.IsInReadQue(pageUrl, models.ReadUrls) {
			//2、提取页面中跟book 有关的url
			booksUrl := models.GetBookUrl(models.GetHtmlStr(pageUrl))
			if len(booksUrl) == 0 {
				continue
			}
			// 判断一下弹出的url 是否已经访问过，如果没有则进行book url 的提取
			for i := 0; i < len(booksUrl); i++ {
				//提取的url 压入到BookUrls 队列中
				//fmt.Println("bookurl",booksUrl[i])
				models.PutInQue(models.BookUrls, booksUrl[i])
			}
		}
		//3、把弹出的url 放入已经访问过的url队列中
		models.PutInQue(models.ReadUrls, pageUrl)

		time.Sleep(time.Second * 2)
	}

	for i := 0; i <= 2; i++ {
		subUrl := models.GetUrl(models.BookUrls)
		if subUrl == "" {
			break
		}
		wg.Add(1)
		go func() {
			if !models.IsInReadQue(subUrl, models.ReadUrls) {
				book.Name = models.GetBookName(models.GetHtmlStr(subUrl))
				book.Score = models.GetBookScore(models.GetHtmlStr(subUrl))
				book.Author = models.GetBookAuthor(models.GetHtmlStr(subUrl))
				book.Isbn = models.GetBookIsbn(models.GetHtmlStr(subUrl))
			}
			models.PutInQue(models.ReadUrls, subUrl)
			c.Ctx.WriteString(fmt.Sprintf("书名：%s\t\t|\t作者：%s\t\t|\t评分：%f\t\t|\tISBN：%d\t\t\n ", book.Name, book.Author, book.Score, book.Isbn))
			wg.Done()
		}()
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
}
