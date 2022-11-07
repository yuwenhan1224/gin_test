package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//静态文件 :html页面用到的样式文件 .css js 文件 图片

func main(){
	r := gin.Default() //创建一个默认路由引擎
	//加载静态文件，
	//"/xxx"的意思是给所有的静态文件起一个统一的路径名     "./lesson09/statics"是实际存放静态文件的文件夹路径
	//（1）-------------------------------------定义模板----------------------------------------------
	r.Static("/xxx","./lesson09/statics")
	//（2）-------------------------------------模板解析----------------------------------------------
	//r.LoadHTMLFiles("./lesson09/templates/posts/index.tmpl","./lesson09/templates/users/index.tmpl")//模板解析
	//gin框架中给模板添加自定义函数模板
	r.SetFuncMap(template.FuncMap{
		"safe":func(str string)template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("./lesson09/templates/**/*")//这样用正则写方便,//模板解析
	//（3）-------------------------------------渲染模板----------------------------------------------
	r.GET(" /posts/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{      //模板渲染
			"title":"liwenzhou",
		})//也可以直接写200
	})
	r.GET("/users/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{      //模板渲染
			"title":"<a href='https://liwenzhou.com'>李文周的博客</a>",
		})//也可以直接写200
	})
	//返回从网上下载的模板
	{{/*
		没有通过define给模板文件起名字的话，模板名字默认是自己的文件名
		*/}}
	r.GET("/home",func(c *gin.Context){
		c.HTML(http.StatusOK,"home.html",nil)
	})
	err:=r.Run(":8080")                               //启动server
	if err !=nil{
		fmt.Println("Http server 启动失败,err:%v",err)
	}

}
