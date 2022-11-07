package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//   1.定义模型
type User struct {
    gorm.Model //CreatedAt UpdatedAt DeletedAt
	Name string
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

	//2. 把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	//u1:=User{Name:"yuwenhan",Age:23}
	//db.Debug().Create(&u1)    //在数据库中创建了一条yuwenhan的记录//.Debug可以被sql语句打印出来
	//u2:=User{Name:"wangyuanyaun",Age:20}
	//db.Debug().Create(&u2)

	//4. 查询
    //var user User  //声明模型结构体类型变量user
	user :=new(User)   //new一般是基本数据类型，new什么类型就返回什么类型的指针，make一般用于map slice chanel
	db.First(&user) //在go语言里面函数和方法的参数的传递都是值拷贝，所以想要真正改变user的值只能使用指针
	fmt.Printf("user:%#v\n",user)

	var users []User
	db.Debug().Find(&users)
	fmt.Printf("users:%v\n",users)






}













