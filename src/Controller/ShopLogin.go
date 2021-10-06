package controller

import (
	"github.com/gin-gonic/gin"
	// "empolder/src/model"
	"empolder/src/config"
)

func ShopLogin(c *gin.Context, code string)  {
	flag := config.CaptchaVerify(c, code)
	if !flag {
		c.JSON(200, config.CodeJSON(201, "验证码不正确", nil))
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "admin" && password == "123" {
			c.JSON(200, config.CodeJSON(200, "登录成功", nil))
		} else {
			c.JSON(200, config.CodeJSON(201, "用户或者密码错误", nil))
		}
	}
}