package model

import (
	"backend/dao"
)

type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func GetTodoList() (todoList []Todo,err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil,err
	}
	return todoList,nil
}

func CreatTodo(todo *Todo) (err error){
	err = dao.DB.Create(todo).Error
	if err != nil {
		return err
	}
	return
}
func GetOne(id uint) (todo *Todo,err error){
	//切记，指针要new出空间之后，查完数据库后才可以向里面放东西
	todo=new(Todo)

	//数据库中找到原先的
	err = dao.DB.Where("id=?",id).First(todo).Error
	if err != nil {
		return nil,err
	}
	return

}

func UpdateTodo(todo *Todo) (err error){
	//保存更改
	err = dao.DB.Save(todo).Error

	if err != nil {
		return err
	}
	return
}


func DeleteTodo(id uint) (err error){
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return
}
