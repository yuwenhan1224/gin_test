package main

import "github.com/gin-gonic/gin"
//---------------------------------------------中间件----------------------------------------------
func main(){
   r:=gin.Default()  //默认使用Logger和Recovery中间件

   r.GET("/index",func(c *gin.Context){

   })

}