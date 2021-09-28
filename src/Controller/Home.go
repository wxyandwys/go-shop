package controller

import (
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context)  {
	c.JSON(200, "{a:33,b:gg}")
}