package controller

import (
	"gin_test/lesson26_bubble/bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"

)
/*
 url   -->  control-->logic-->models
请求来了-->  控制器 --> 业务逻辑 --> 模型层的增删改查

*/

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
}
func CreateATodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 会发请求到这里

	//1. 从请求中把数据拿出来

	var todo models.Todo //前端访问/v1/todo这个路由，通过点击emit按钮,发送post请求给服务端，现在前端的请求发送到了这里，定义一个todo变量去接收数据
	c.BindJSON(&todo)    //把请求中的参数与我们后端的结构体进行绑定吗，如果请求中有了	ID Title Status 这三个字段，bindjson会自动帮我们取出值并帮我们初始化机构体
	//bindjson不需要自己定义状态码
	//2. 存入数据库
	err:= models.CreateAToDo(&todo)
	if err !=nil{
		//创建记录失败
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		//创建记录成功
		//c.JSON(http.StatusOK,todo)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,      //定义错误码
			"msg":  "请求成功success", //提示
			"data": todo,      //数据
		})
	}
}
func GetTodoList(c *gin.Context) {
	//查询todos 这个表里所有的数据
	todoList,err:= models.GetAllTodo()
	 if err!=nil{
		c.JSON(http.StatusOK, err.Error()) //查询失败
	} else {
		c.JSON(http.StatusOK, todoList) //查询成功，通过把todoList切片打包成json格式是数据来给前端返回数据
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id") //获取URL里面的id参数
	if !ok {                     //
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
		return //加个return终止代码的运行
	}
      todo,err:=models.GetATodo(id)
	  if err!=nil{
		  c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		  return
	  }
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil { //Save()默认会更新该对象的"所有"字段，即使你没有赋值
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	//获取请求参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	if err := models.DeleteATodo(id);err !=nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "delete"})
	}

}