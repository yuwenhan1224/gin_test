package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./lesson15/index.html")

	r.GET("/index",func(c *gin.Context){
       c.HTML(http.StatusOK,"index.html",nil)
	})
	//处理multipart forms提交文件时默认的内存限制是32 MiB
	r.POST("/upload",func(c *gin.Context){
		//从请求中中去读取文件
		f,err:=c.FormFile("f1")//从请求中获取携带的参数一样
		if err !=nil{
			c.JSON(200,gin.H{
				"error":err.Error(),
			})
		}else{
			//将读取的文件保存在本地（服务端本地）
			//dst:=fmt.Sprintf("./lesson15/%s",f.Filename)
			dst2:=path.Join("./lesson15",f.Filename)
			c.SaveUploadedFile(f,dst2)
		}
		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
		})
	})
	r.Run(":9000")
}
