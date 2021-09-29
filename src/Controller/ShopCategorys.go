package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
)

// 获取类型数据
func GetShopCategoryList(c *gin.Context)  {
	categorys := []model.ShopCategorys{}
	config.DBHelper.Find(&categorys)
	c.JSON(200, gin.H{"code":200, "message": "成功", "data":categorys})
}