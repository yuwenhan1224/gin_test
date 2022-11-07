package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//shouldbind可以根据前端请求的数据类型（queryString、form表单、json数据）去自动获取对应的数据
//结构体想要从外面被别人通过反射去取到字段，要大写，还  要加上tag
//target``把url里的参数和结构体里的参数对应起来
type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./lesson14/index.html")
	r.GET("/user",func(c *gin.Context){
		//获取用户请求时候发来的请求数据，方法一:
		//username1:=c.Query("username")//把浏览器输入的username=233&passwrd=555的值赋值给username1和password2
		//password2:=c.Query("password")

		//u:=UserInfo{
		//	username: username1,
		//	password: password2,
		//}
		//fmt.Printf("%#v\n",u)
		//c.JSON(http.StatusOK,gin.H{
		//	"message":"ok",
		//})
		//获取用户请求时候发来的请求数据，方法二:
		var u UserInfo    //声明一个UserInfor类型变量u
		//如果请求里面出现了username和password这两个字段，就把请求里面的username，passwrd的值取出来放到u里面
		//c.ShouldBind(&u)内部通过反射去找到结构体里找有几个字段
		//因为go语言的函数和方法的参数是值拷贝，想要改变原来的值只能使用指针
		err:=c.ShouldBind(&u) //把请求的数据绑定到u

		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}

		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"username": "username",
			"password": "password",

		})

	})
	//用户访问index这个页面时候，给用户返回index.html这个页面
	r.GET("/index",func(c *gin.Context){
		c.HTML(200,"index.html",nil)
	})

	//form参数绑定
	r.POST("/form",func(c *gin.Context){
		var u UserInfo    //声明一个UserInfor类型变量u
		err:=c.ShouldBind(&u) //因为go语言的函数和方法的参数是值拷贝，想要改变原来的值只能使用指针
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)//把结构体实例里面的值打印出来
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
  //json格式参数绑定
	r.POST("/json",func(c *gin.Context){
		var u UserInfo    //声明一个UserInfor类型变量u
		err:=c.ShouldBind(&u) //因为go语言的函数和方法的参数是值拷贝，想要改变原来的值只能使用指针
		if err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)//把结构体实例里面的值打印出来
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})



	r.Run(":9000")
}