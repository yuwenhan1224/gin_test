package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//   1.定义模型
type User struct {
	ID int64
	//Name string `gorm:"default:'小王子'"`
	//（1）指针方式将零值存入数据库
	//Name *string `gorm:"default:'小王子'"`
	//（2）使用Scanner/valuer存入数据库
	Name sql.NullString `gorm:"default:'小王子'"`
	Age int64
}
func main() {

	//连接数据库
	db, err := gorm.Open("mysql", "root:y760069562@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local") //parseTime=True&loc=Local"解析本地时间
	//                     "要连接的数据库名","用户名:密码@(连接的数据库所在的主机地址)/所要连接的数据库名?编码规则&解析时间类型&本地数据"
	if err != nil {
		fmt.Println(err) //如果连接不上就抛出错误，跳出程序
	}
	defer db.Close() //程序运行结束后把数据连接关闭

	//2. 把模型和数据库中发表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	//u:=User{Name:"yuwenhan",Age:23}
	//u:=User{Age:23}
	//（1）指针方式将零值存入数据库
	//u:=User{Name:new(string),Age:23}

	//（2）使用Scanner/valuer存入数据库
	u:=User{Name:sql.NullString{String:"",Valid: true},Age:23}
	db.Delete(&u)
	db.NewRecord(&u)  //判断主键是否为空
	db.Debug().Create(&u)    //在数据库中创建了一条yuwenhan的记录//.Debug可以被sql语句打印出来
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空

}