package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
//	"log"
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

func GetShopTree(c *gin.Context)  {
	trees := []model.ShopTrees{}
	shopId := c.Query("id")
	config.DBHelper.Where("shops_id = ?", shopId).Find(&trees)
	for i, data := range trees {
		treeVs := []model.ShopTreeVs{}
		config.DBHelper.Where("tree_id = ?", data.Id).Find(&treeVs)		
		trees[i].Vs = treeVs
	}
	data :=  map[string]interface{}{}
	data["trees"] = trees
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

func GetShopListSku(c *gin.Context)  {
	ls := []model.ShopLists{}
	shopId := c.Query("id")
	config.DBHelper.Where("shop_id = ?", shopId).Find(&ls)
	for i, data := range ls {
		sls := []model.ShopListSkus{}
		//config.DBHelper.Where("id = ?", data.Id).Find(&sls)
		config.DBHelper.Table("shop_list_skus").Select("shop_list_skus.id,shop_list_skus.k,shop_list_skus.v,ks").Joins("left join shop_trees on shop_trees.id = shop_list_skus.k").Where("shop_list_skus.id = ?", data.Id).Scan(&sls)
		ls[i].Slsu = sls
	}
	data := map[string]interface{}{}
	data["list"] = ls
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

func GetShopParameters(c *gin.Context)  {
	parameters := []model.ShopParameters{}
	shopId := c.Query("id")
	config.DBHelper.Where("shop_id = ?", shopId).Find(&parameters)
	data := map[string]interface{}{}
	data["parameters"] = parameters
	c.JSON(200, config.CodeJSON(200, "成功", data))
}