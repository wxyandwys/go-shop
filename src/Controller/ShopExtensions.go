package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
)

func GetShopExtensionList(c *gin.Context)  {
	extension := []model.ShopExtensions{}
	config.DBHelper.Find(&extension)
	c.JSON(200, gin.H{"code": 200, "message": "成功", "data": extension})
}
