package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model      //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64   `gorm:"column:beast_id"`  //零值类型的age
	Birthday     *time.Time
	//结构体的Tag设置
	Email        string  `gorm:"type:varchar(100);unique_index"`//unique_index规定Email为唯一的索引，不可重复
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
type Animal struct {
	AnimalId    int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
	Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
	Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}
//自己设置数据库名 唯一指定表名
func (Animal)TableName()string{
	return  "yuwenhan" ///返回自己设置的数据库名
}
func main() {
    //可以给表明添加默认的表名前缀后缀等
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "prefix_" + defaultTableName;
	}

	//连接数据库
	db, err := gorm.Open("mysql", "root:y760069562@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local") //parseTime=True&loc=Local"解析本地时间
	//                     "要连接的数据库名","用户名:密码@(连接的数据库所在的主机地址)/所要连接的数据库名?编码规则&解析时间类型&本地数据"
	if err != nil {
		fmt.Println(err) //如果连接不上就抛出错误，跳出程序
	}
	defer db.Close() //程序运行结束后把数据连接关闭

	db.SingularTable(true)//// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//创建表 自动迁移 （把结构体和数据表进行对应）
	 db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
    //使用User结构体 建叫做xiaownagzi 的表
	db.Table("xiaowangzi").CreateTable(&User{})


}