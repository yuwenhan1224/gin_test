package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"hello golang",
	})
}
func main() {
	r := gin.Default() //创建一个默认路由引擎
	//指定用户使用GET请求返访问/hellO时，执行sayHello这个函数
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"method":"GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})

	//启动服务
	r.Run()
}