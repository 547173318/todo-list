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
	ID uint `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
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

	//创建一个初始化的gin引擎
	r := gin.Default()


	//配置路由
	userGroup := r.Group("/v1")
	{
		userGroup.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			db.Find(&todos)
			c.JSON(http.StatusOK,todos)

		})
		userGroup.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.ShouldBind(&todo)
			db.Create(&todo)
			c.JSON(http.StatusOK,todo)
		})
		userGroup.PUT("/todo/:id", func(c *gin.Context) {
			var todo Todo
			id,_:=c.Params.Get("id")

			//数据库中找到原先的
			db.Where("id=?",id).Find(&todo)

			//更改数据库中找到的
			c.ShouldBind(&todo)

			//保存更改
			db.Save(&todo)
		})
		userGroup.DELETE("/todo/:id", func(c *gin.Context) {
			id,_:=c.Params.Get("id")
			db.Where("id=?",id).Delete(&Todo{})
		})

	}

	//启动服务，默认在port8080
	r.Run(":9000")
}