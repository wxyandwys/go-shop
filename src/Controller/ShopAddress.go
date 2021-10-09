package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
)
// 添加地址
func InsertAddress(c *gin.Context)  {
	address := model.ShopAddresses{}
	c.Bind(&address)
	u := config.ShowUser(c)
	address.UserId = u.Id

	tmp := config.DBHelper.Create(&address)
	data := map[string]interface{}{}
	data["index"] = tmp
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

// 查询地址
func ShowAddress(c *gin.Context)  {
	u := config.ShowUser(c)
	address := []model.ShopAddresses{}
	config.DBHelper.Where("user_id = ?", u.Id).Find(&address)
	data := map[string]interface{}{}
	data["address"] = address
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

// 修改地址
func UpdateAddress(c *gin.Context)  {
	address := model.ShopAddresses{}
	c.Bind(&address)
	tmp := config.DBHelper.Save(&address)
	data := map[string]interface{}{}
	data["index"] = tmp
	c.JSON(200, config.CodeJSON(200, "成功", data))
}

// 查询地址id
func ShowAddressById(c *gin.Context)  {
	address := model.ShopAddresses{}
	u := config.ShowUser(c)
	addressId := c.Query("id")
	config.DBHelper.Where("user_id = ? and id = ?", u.Id, addressId).First(&address)
	data := map[string]interface{}{}
	data["address"] = address
	c.JSON(200, config.CodeJSON(200, "成功", data))
}