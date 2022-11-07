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
	Active bool
}
func main() {

	//2.连接数据库
	db, err := gorm.Open("mysql", "root:y760069562@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local") //parseTime=True&loc=Local"解析本地时间
	//                     "要连接的数据库名","用户名:密码@(连接的数据库所在的主机地址)/所要连接的数据库名?编码规则&解析时间类型&本地数据"
	if err != nil {
		fmt.Println(err) //如果连接不上就抛出错误，跳出程序
	}
	defer db.Close() //程序运行结束后把数据连接关闭

	//3. 把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 4.创建
	u1:=User{Name:"yuwenhan",Age:23,Active: true}
	db.Debug().Create(&u1)    //在数据库中创建了一条yuwenhan的记录//.Debug可以被sql语句打印出来
	u2:=User{Name:"wangyuanyaun",Age:20,Active: false}
	db.Debug().Create(&u2)

	//5. 查询
    var user User  //声明模型结构体类型变量user
	//user :=new(User)   //new一般是基本数据类型，new什么类型就返回什么类型的指针，make一般用于map slice chanel
	db.First(&user) //在go语言里面函数和方法的参数的传递都是值拷贝，所以想要真正改变user的值只能使用指针
	fmt.Printf("user:%#v\n",user)

	//6. save更新全部，update更新单个属性，uodates更新多个属性
	user.Name="qimi"
	user.Age=99
	db.Debug().Save(&user)  //默认会修改所有字段
	db.Debug().Model(&user).Update("name","小王子")

	m1:=map[string]interface{}{
		"name":"yyyy",
		"age": "18",
		"active":true,
	}
	db.Debug().Model(&user).Updates(m1) //m1列出来都全部更新
	db.Debug().Model(&user).Select("age").Updates(m1) //只更新age字段
	db.Debug().Model(&user).Omit("active").Updates(m1) //除了active字段都更新

	//让users表中所有的用用户的年龄在原来的基础上+2岁
	db.Model(&User{}).Update("age",gorm.Expr("age+?",2))



}













