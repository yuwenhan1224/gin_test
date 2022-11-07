package main
//重定向的两种方式
import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main(){
  r:=gin.Default()
  //使用Redirect进行HTTP重定向
  r.GET("/index",func(c * gin.Context){
    //c.JSON(200,gin.H{
    //  "status":"ok",
    //})
    c.Redirect(http.StatusMovedPermanently,"htttp://sogo.com")
  })
  //使用HandleContext()进行路由重定向
  r.GET("/a",func(c * gin.Context){
    //跳转到 /b 对应的路由处理函数
    c.Request.URL.Path="/b"  // 把请求的URL修改
    r.HandleContext(c)       //
  })
    r.GET("/b",func(c * gin.Context){
      c.JSON(http.StatusOK,gin.H{
        "message":"b",
      })
    })

  r.Run(":9000")
}
