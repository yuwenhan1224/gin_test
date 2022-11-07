package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	//访问/index的GET请求会走一条逻辑处理
	//路由
	//GET:获取信息
	r.GET("/index",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})
	//post一般用于form表单的提交，包括一些上传文件、填用户名密码等
	//如注册会员，把用户名密码发送给服务器，服务器再把用户名密码的信息存到数据库里面
	r.POST("/index",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	//PUT修改个人的部分信息，局部更新
	r.PUT("/index",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	//DELETE 删除信息
	r.DELETE("/index",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method":"delete",
		})
	})
	//请求方法大杂烩
	r.Any("/user",func(c *gin.Context){
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK,gin.H{"method":"any"})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{"method":"post"})
       // ....
		}
	})

	//NoRoute,当用户请求的URL是定义的其他URL时，执行下面的函数
	r.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound,gin.H{
			"message":"NoRoute5555",
		})
	})

	//视频的首页和详情页
	//r.GET("video/index",func(c *gin.Context){
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	//})
	//r.GET("video/xx",func(c *gin.Context){
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	//})
	//r.GET("shop/index",func(c *gin.Context){
	//	c.JSON(http.StatusOK,gin.H{"msg":"/shop/index"})
	//})

	//路由组的组 多用于区分不同业务线或者API版本
	//把公用的前缀提取出来，创建一个路由组
	//路由是支持嵌套的
	videGroup:=r.Group("/video")
	{
		videGroup.GET("/index",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
		})
		videGroup.GET("/xx",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{"msg":"/video/xx"})
		})
		videGroup.GET("/oo",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{"msg":"/video/oo"})
		})
	}










	r.Run(":9000")
}
