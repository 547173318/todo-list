package main

import (
	"backend/dao"
	"backend/model"
	"backend/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)




func main()  {
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	
	dao.DB.AutoMigrate(&model.Todo{})

	//创建一个初始化的gin引擎,并且配置路由
	r :=router.SetUpRouter()

	//启动服务，默认在port8080
	r.Run(":9000")
}