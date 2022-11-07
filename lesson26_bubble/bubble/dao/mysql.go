package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //数据库的驱动
)


//定义数据库全局变量
var(
	DB *gorm.DB
)



//连接数据库
func InitMySQL()(err error){

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

func Close(){
	DB.Close()
}