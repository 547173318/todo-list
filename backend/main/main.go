package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var(
	db *gorm.DB
)

//mode
type Todo struct {
	gorm.Model
	Title string
	Status bool
}

//数据库连接
func InitMySql()(err error){
	dist := "root:123456@(127.0.0.1:3306)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql",dist)
	if err != nil{
		return
	}
	return db.DB().Ping()
}


func main()  {
	err := InitMySql()
	if err != nil {
		panic(err)
	}
	defer db.Close()


	db.AutoMigrate(&Todo{})

	//var todo Todo{}
	//db.Create(todo)



	//创建一个初始化的gin引擎
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK,"index")
	})

	//启动服务，默认在port8080
	r.Run()
}