package config

/*
按照MVC原理
把路由抽象出来，不要写到MAIN文件里面
**/
import (
	"gingorm/api"

	"github.com/gin-gonic/gin"
)

// 在main里面不用写 r*gin.Engine。在这里需要写
func Routes(r *gin.Engine) {
	//从Main里面调过来后，再将具体方法写到ping里面更清爽。路由为注册
	r.GET("/ping", api.Ping)
}
