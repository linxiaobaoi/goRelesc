package router

import (
	"github.com/gin-gonic/gin"
	"github.com/linxiaobaoi/goRelesc.git/controller"
)

func AppRouter() {
	r := gin.Default()
	//用户路由
	user := r.Group("/user")
	{
		//发送验证码
		user.POST("/send", controller.SendCms)
	}
	r.Run()
}
