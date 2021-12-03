package dao

import "github.com/jinzhu/gorm"

var(
	DB *gorm.DB
)



//数据库连接
func InitMySql()(err error){
	dist := "root:123456@(127.0.0.1:3306)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dist)
	if err != nil{
		return
	}
	return DB.DB().Ping()
}

