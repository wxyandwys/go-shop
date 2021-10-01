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

func GetShop(c *gin.Context)  {
	id := c.Query("id")
	shop := model.Shops{}
	imgs := []model.ShopImgs{}
	config.DBHelper.Where("id = ?", id).First(&shop)
	config.DBHelper.Where("id = ?", id).Find(&imgs)
	data :=  map[string]interface{}{}
	data["shop"] = shop
	data["imgs"] = imgs
	// c.JSON(200, gin.H{"code": 200, "message": "成功", "data": data })
	c.JSON(200, config.CodeJSON(200, "成功", data))
}