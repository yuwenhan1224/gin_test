package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main(){
	r:=gin.Default()//默认路由
	//func(c *gin.Context)是指定的函数
	r.GET("/json",func(c *gin.Context){
		//方法1：使用map
		//data:=map[string]interface{}{
		//	"name":"汪圆圆",
		//	"message":"hello",
		//	"age":18,
		//}
		//gin.H是gin定义好的一个map的快捷方式，方便进行json数据的传输
		data:=gin.H{"name":"ddd","message":"hello","age":18}
		c.JSON(http.StatusOK,data)
	})
	//方法二：结构体 灵活使用tag来对结构体字段做定制化操作
	//使用结构体类型去承载数据，最后通过Json包对结构体实例进行json序列化的操作，最后发送给前端
	type msg struct{
			Name string `json:"name"`
			Message string `bson:"xxx"`
			Age int
		}
	r.GET("/json2",func(c *gin.Context){
		   //结构体里的key首字母必须大写
			data:=msg{
				Name:"汪圆圆2",
				Message:"hello2",
				Age:182,
			}
		c.JSON(http.StatusOK,data)//json的序列化
	})


	err:=r.Run(":9000")
	if err!=nil{
		fmt.Println("http run failure,err:%v",err)
	}
}
