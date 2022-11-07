package models

import (
	"gin_test/lesson26_bubble/bubble/dao"
)

//Todo Model
type Todo struct {
	ID int `json:"id"`//因为和前端使用json格式的数据，所以用json的tag
	Title string `json:"title"`//
	Status bool `json:"status"`//
}
/*

Todo 这个Model的增删改查的操作都放在这里
 */

//创建todo
func CreateAToDo(todo *Todo)(err error){
	 err = dao.DB.Create(&todo).Error
     return
}

func GetAllTodo()(todolist [] *Todo,err error){
	if err = dao.DB.Find(&todolist).Error; err != nil{
		return nil,err
	}
	return
}

func GetATodo(id string)(todo *Todo,err error){
   todo =new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil,err
	}
	return
}
func UpdateATodo(todo *Todo)(err error){
	 err=dao.DB.Save(todo).Error
	 return
}
func DeleteATodo(id string)(err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return

}