package main

/*
gin通俗来说就两个功能 1注册路由 2 返回JSON的数据 还有可能就中间件
gorm 访问数据库


**/
import (
	"GoGinGorm/config"
	"GoGinGorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//要在路由前调用数据库的链接
	config.Init()
	routes.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
