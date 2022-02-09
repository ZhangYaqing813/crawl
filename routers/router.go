package routers

import (
	"crawl_books/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/books", &controllers.BooksController{},"*:BooksInfo")
}
