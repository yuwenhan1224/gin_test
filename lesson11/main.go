package main
//querystring
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	//querystring

	//GET请求 URL ?后面的是querystring参数
	//key=value格式，多个key-value用 & 连接
	//eq: /web/query=小王子&age=18
	//跟请求相关的都要在c里面找
	r.GET("/web",func( c *gin.Context){
		//http的get请求的请求头最好不要写太多
		//获取浏览器那边发请求携带的query string 参数的三种方式
		name1:=c.Query("query")    //通过Query获取请求中携带的querystring 参数
		name2:=c.DefaultQuery("query","somebody")//取不到就用默认的值
		name3,ok:=c.GetQuery("query")        //取到返回(值，true),取不到返回("",false)
		age:= c.Query("age")
		if !ok{
			//取不到
			name3="somebody"
		}
		//当浏览器访问/web时候，执行func( c *gin.Context)这个函数，而c.JSON(http.StatusOK,gin.H{则可以给页面返回name1,name2..等数据
     c.JSON(http.StatusOK,gin.H{

		 "name1":name1,
		 "name2":name2,
		 "name3":name3,
		 "age":age,
	 })
	})

	err:=r.Run(":8080")
	if err!=nil{
		fmt.Println("http run failure,err:%v",err)
	}
}