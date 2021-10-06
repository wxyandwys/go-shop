package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
	"crypto/md5"
	"fmt"
	"time"
)

func ShopLogin(c *gin.Context, code string)  {
	flag := config.CaptchaVerify(c, code)
	if !flag {
		c.JSON(200, config.CodeJSON(201, "验证码不正确", nil))
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		users := model.ShopUsers{}
		config.DBHelper.Where("username = ?", username).First(&users)
		psw := md5.Sum([]byte(password + users.Sale)) // md5 + sale
		pswMd5 := fmt.Sprintf("%x", psw)
		if username == users.Username && pswMd5 == users.Password {
			userToken := md5.Sum([]byte(username + fmt.Sprintf("%x", time.Now().Unix()) ))
			conn := config.RedisDefaultPool.Get()
			conn.Do("setex", fmt.Sprintf("%x", userToken), 60 * 60 , users)
			data := map[string]interface{}{}
			data["userToken"] = fmt.Sprintf("%x", userToken)
			c.JSON(200, config.CodeJSON(200, "登录成功", data))
		} else {
			c.JSON(200, config.CodeJSON(201, "用户或者密码错误", nil))
		}
	}
}