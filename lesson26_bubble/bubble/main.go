package main

import (
	"fmt"
	"gin_test/lesson26_bubble/bubble/dao"
	"gin_test/lesson26_bubble/bubble/models"
	"gin_test/lesson26_bubble/bubble/routers"

)

func main(){
	//手动创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
     err:=dao.InitMySQL()
	if err != nil {
		//panic(err)
		fmt.Println(err) //如果连接不上就抛出错误，跳出程序
	}
	defer dao.Close()//程序运行结束后把数据连接关闭
    //模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //数据库中的表名应该是todos
   r:=routers.SetupRouter()
	r.Run(":8080")
}
