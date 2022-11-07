
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取form表单提交的参数

func main(){
	r:=gin.Default()
	//（1）---------------------------------------定义模板------------------------
	r.LoadHTMLFiles("./lesson12/login.html","./lesson12/index.html")//加载模板
    //r.GET只能处理login的GET请求，不能处理登录按钮的post请求
	r.GET("/login",func(c *gin.Context){    //当用户访问/login时候，通过func函数返回给前端一个模板文件
		// 没有通过define给模板文件起名字的话，模板名字默认是自己的文件名，现在的文件名是login.html
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//（2）---------------------------------------解析模板------------------------
	//一次请求对应一次响应，跟请求相关的都要在c里面找
    //login post
	r.POST("/login",func(c *gin.Context){
		//获取form表单提交的数据
		//方法一:
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		//方法二：
		username2:=c.DefaultPostForm("username","sombody")
		password2:=c.DefaultPostForm("password","****")
		//（3）-----------------------------------渲染模板------------------------
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
			"Name2":username2,
			"Password2":password2,
		})

	})

	err:=r.Run(":9000")
	if err!=nil{
		fmt.Println("http run failure,err:%v",err)
	}
}

