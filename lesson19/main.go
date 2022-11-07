package main

//当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//UserInfor -->数据表
type  UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main(){
	//连接数据库
	db,err:=gorm.Open("mysql","root:y760069562@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")//parseTime=True&loc=Local"解析本地时间
	 //                     "要连接的数据库名","用户名:密码@(连接的数据库所在的主机地址)/所要连接的数据库名?编码规则&解析时间类型&本地数据"
	if err !=nil{
		fmt.Println(err)//如果连接不上就抛出错误，跳出程序
	}
	defer db.Close()//程序运行结束后把数据连接关闭

	//创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1:=UserInfo{1,"fear","ff","ggg"}
	db.Create(&u1)
	fmt.Println(u1)
	//查询
	var u UserInfo
	db.First(&u)//只能传入指针才能修改变量,查询表中一天的数据保存到u
	fmt.Printf("u:%#v\n",u)
	//更新
	db.Model(&u).Update("hobby","双色球")
	fmt.Printf("u:%#v\n",u)
    //删除
	db.Delete(&u)
	fmt.Printf("u:%#v\n",u)
	//创建数据行
	u2:=UserInfo{1,"余文汉","男","唱跳Rap篮球"}
	db.Create(&u2)
	fmt.Printf("u2:%#v\n",2)

}