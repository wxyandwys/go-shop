package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
)

// 推送获取类型数据
func GetShopCategoryList(c *gin.Context)  {
	categorys := []model.ShopCategorys{}
	config.DBHelper.Limit(8).Find(&categorys)
	c.JSON(200, gin.H{"code":200, "message": "成功", "data":categorys})
}

//  获取所有父分类
func GetShopCategoryListParent(c *gin.Context)  {
	categorys := []model.ShopCategorys{}
	config.DBHelper.Where("pid = ?", 0).Find(&categorys)
	data := map[string]interface{}{}
	data["categoryParent"] = categorys
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

//  获取所有子分类
func GetShopCategoryListChildren(c *gin.Context)  {
	categoryId := c.Query("id")
	categorys := []model.ShopCategorys{}
	config.DBHelper.Where("pid = ?", categoryId).Find(&categorys)
	data := map[string]interface{}{}
	data["categoryChildren"] = categorys
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

func GetShopCategoryListChildrenById(c *gin.Context)  {
	id := c.Query("id")
	shops := []model.Shops{}
	config.DBHelper.Where("category_id = ?", id).Find(&shops)
	data := map[string]interface{}{}
	data["shops"] = shops
	c.JSON(200, config.CodeJSON(200, "成功", data))
}