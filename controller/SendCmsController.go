/**
* Created by GoLand
* User: lingm
* Date: 2023/6/12
* Time: 下午 03:09
* Author: 现在的努力是为了小时候吹过的NB
* Atom: 小白从不写注释！！！
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/linxiaobaoi/goRelesc/config"
	"github.com/linxiaobaoi/goRelesc/service"
	"net/http"
)

// 发送验证码
func SendCms(c *gin.Context) {
	email := c.PostForm("email")
	str, err := service.Send(email)
	if err == false {
		data := config.ReturnErrorMessage(0, str)
		c.JSON(http.StatusOK, data)
		return
	} else {
		data := config.ReturnMessage(1, "发送成功", str)
		c.JSON(http.StatusOK, data)
		return
	}

}
