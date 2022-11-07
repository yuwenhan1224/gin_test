package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	// 获取请求path URL 参数
	// 注意URL的匹配不要冲突
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month/:day", func(c *gin.Context) {
		// 获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		day := c.Param("day")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
			"day":   day,
		})
	})
	r.Run(":9999")
}
//package lesson13
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func fff() {
//	r := gin.Default()
//	//http://localhost:9000/小王子/18
//
//	//渲染数据方法一:---------------------------------------
//	r.GET("/:name/:age", func(c *gin.Context) {
//		//获取路径参数
//		//name=小王子，age=18
//		name := c.Param("name")
//		age := c.Param("age")
//		c.JSON(http.StatusOK, gin.H{
//			"name": name,
//			"age":  age,
//		})
//	})
//	////渲染数据方法二:-----------------------------------------
//	//r.GET("/:name2/:age2", func(c *gin.Context) {
//	//	//获取路径参数
//	//	//name=小王子，age=18
//	//	name2 := c.Param("name2")
//	//	age2 := c.Param("age2")
//	//	data := map[string]interface{}{
//	//		"name2": name2,
//	//		"age2":  age2,
//	//	}
//	//	c.JSON(200, data)
//	//})
//
//	r.Run(":9000")
//}