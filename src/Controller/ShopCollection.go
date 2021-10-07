package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
	// "encoding/json"
	// "github.com/garyburd/redigo/redis"
)

// 查询收藏
func ShowCollection(c *gin.Context)  {
	u := config.ShowUser(c)

	collection := []model.ShopCollections{}
	config.DBHelper.Where("user_id = ?", u.Id).Find(&collection)
	shops := []model.Shops{}

	for _, coll := range collection{
		shop := model.Shops{}
		config.DBHelper.Where("id = ?",coll.ShopId).First(&shop)
		shops = append(shops, shop)
	}

	data := map[string]interface{}{}
	data["shops"] = shops
	c.JSON(200, config.CodeJSON(200, "成功", data))
	
}

// 查询用户是否收藏商品
func ShowCollectionById(c *gin.Context)  {
	shop_id := c.Query("shop_id")
	u := config.ShowUser(c)
	colection := model.ShopCollections{}

	var count int 

	config.DBHelper.Where("user_id = ?", u.Id).Where("shop_id = ?", shop_id).Find(&colection).Count(&count)
	data := map[string]interface{}{}
	data["count"] = count
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

// 添加收藏
func InsertCollection(c *gin.Context)  {
	collection := model.ShopCollections{}
	c.Bind(&collection)

	u := config.ShowUser(c)

	collection.UserId = u.Id
	config.DBHelper.Create(&collection)

	c.JSON(200, config.CodeJSON(200, "成功", nil))
}

// 删除收藏
func DeleteCollection(c *gin.Context)  {
	shopId := c.PostForm("id")
	u := config.ShowUser(c)

	config.DBHelper.Where("shop_id = ? and user_id = ?", shopId, u.Id).Delete(model.ShopCollections{})

	c.JSON(200, config.CodeJSON(200, "成功", nil))	
}