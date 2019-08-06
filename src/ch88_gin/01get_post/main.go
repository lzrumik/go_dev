package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
http://127.0.0.1:8080/ping
http://127.0.0.1:8080/user/ran
http://127.0.0.1:8080/user/john/
http://127.0.0.1:8080/user/john/send
http://127.0.0.1:8080/welcome?firstname=Jane&lastname=Doe
 */
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// 获取post参数  curl -d "" http://127.0.0.1:8080/form_post 必须是post请求 否则404
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// curl -d "name=john&message=123" "http://127.0.0.1:8080/post?id=110&page=20"
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"name":    name,
			"page":    page,
			"id":    id,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}