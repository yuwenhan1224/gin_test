package routers

import (
	"gin_test/lesson26_bubble/bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine{
	r := gin.Default()   //*gin.dDefault生成的就是一个Engine对象
	//加载template下面的所有模板文件,多个文件要用LoadHTMLGlob，单个文件要用loadhtmlfile
	r.LoadHTMLGlob("./lesson25/bubble/templates/*")
	//加载静态文件,第一个参数是路由地址，第二个是静态文件所在位置
	r.Static("/static", "./lesson25/bubble/dist/static")
	//处理函数
	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加,往路由里面添加数据
		v1Group.POST("/todo", controller.CreateATodo)

		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改某一个待办事项 更新操作
		v1Group.PUT("todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("todo/:id", controller.DeleteATodo)
	}
	return  r

}