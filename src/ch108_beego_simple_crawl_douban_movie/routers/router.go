package routers

import (
	"ch108_beego_simple_crawl_douban_movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crawl_movie", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
