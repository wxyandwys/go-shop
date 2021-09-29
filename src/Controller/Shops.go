package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
)

func GetShopList(c *gin.Context)  {
	shops := []model.Shops{}
	config.DBHelper.Find(&shops)
	c.JSON(200, gin.H{"code": 200, "message": "成功", "data": shops})
}