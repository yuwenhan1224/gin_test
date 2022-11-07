package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //数据库的驱动
	"net/http"
)
//定义数据库全局变量
var(
	DB *gorm.DB
)


//Todo Model
type Todo struct {
	ID int `json:"id"`//因为和前端使用json格式的数据，所以用json的tag
   Title string `json:"title"`//
   Status bool `json:"status"`//
}
//连接数据库
func initMySQL()(err error){

	dsn:="root:y760069562@(127.0.0.1)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//"要连接的数据库名","用户名:密码@(连接的数据库所在的主机地址)/所要连接的数据库名?编码规则&解析时间类型&本地数据"
	//下面的 ’DB, err‘的后面没有冒号，因为已经定义了全局变量DB !!!!!!!!!
	DB, err = gorm.Open("mysql",dsn) //parseTime=True&loc=Local"解析本地时间
	if err !=nil{
		return
	}
	err=DB.DB().Ping()//ping的通返回一个nil,ping不通返回一个错误
    return err
}
func main(){
	//手动创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
     err:=initMySQL()
	if err != nil {
		//panic(err)
		fmt.Println(err) //如果连接不上就抛出错误，跳出程序
	}
	defer DB.Close() //程序运行结束后把数据连接关闭
    //模型绑定
	DB.AutoMigrate(&Todo{}) //数据库中的表名应该是todos


	r:=gin.Default()
	//加载template下面的所有模板文件,多个文件要用LoadHTMLGlob，单个文件要用loadhtmlfile
	r.LoadHTMLGlob("./lesson25/bubble/templates/*")
	//加载静态文件,第一个参数是路由地址，第二个是静态文件所在位置
	r.Static("/static","./lesson25/bubble/dist/static")
	//处理函数
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//v1
	v1Group:=r.Group("v1")
	{
		//待办事项
		//添加,往路由里面添加数据
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 会发请求到这里

			//1. 从请求中把数据拿出来

			var todo Todo     //前端访问/v1/todo这个路由，通过点击emit按钮,发送post请求给服务端，现在前端的请求发送到了这里，定义一个todo变量去接收数据
			c.BindJSON(&todo) //把请求中的参数与我们后端的结构体进行绑定吗，如果请求中有了	ID Title Status 这三个字段，bindjson会自动帮我们取出值并帮我们初始化机构体
			//bindjson不需要自己定义状态码
			//2. 存入数据库

			//err=DB.Create(&todo).Error
			//if err !=nil{
			//	fmt.Printf("err:%v",err)
			//}

			//3.返回响应
			//下面是把存入数据和返回响应两个操作放在一起了
			if err = DB.Create(&todo).Error; err != nil {
				//创建记录失败
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				//创建记录成功
				//c.JSON(http.StatusOK,todo)
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,      //定义错误码
					"msg":  "success", //提示
					"data": todo,      //数据
				})
			}

		})

		//查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todos 这个表里所有的数据
			var todoList []Todo //定义一个切片去接受数据库返回的所有数据
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, err.Error()) //查询失败
			} else {
				c.JSON(http.StatusOK, todoList) //查询成功，通过把todoList切片打包成json格式是数据来给前端返回数据
			}
		})

		//查看某一个待办事项,id对应数据库中某条数据的id
		v1Group.GET("/todo/:id", func(c *gin.Context) {
		})

		//修改某一个待办事项 更新操作
		v1Group.PUT("todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id") //获取URL里面的id参数
			if !ok {                     //
				c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
				return //加个return终止代码的运行
			}
			var todo Todo //定义变量接收从数据库返回的值
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return //加个return终止代码的运行
			}

			c.BindJSON(&todo) //把数据库的数据绑定到todo！bindjson状态码不需要自己定义
			// bindjson相当于前端收据与数据库之间的一个中介，如果没有这个中介就得自己手动一个个获取，再一个个存到结构体中

			//更新操作
			if err = DB.Save(&todo).Error; err != nil { //Save()默认会更新该对象的"所有"字段，即使你没有赋值
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除
		v1Group.DELETE("todo/:id", func(c *gin.Context) {
			//获取请求参数
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error;err !=nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "delete"})
			}

		})
	}

	r.Run(":8080")
}
