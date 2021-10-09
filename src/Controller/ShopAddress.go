package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
	"strconv"
)
// 添加地址
func InsertAddress(c *gin.Context)  {
	address := model.ShopAddresses{}
	c.Bind(&address)
	u := config.ShowUser(c)
	address.UserId = u.Id
	if address.IsDefault {
		addresses := []model.ShopAddresses{}
		config.DBHelper.Where("user_id = ?", u.Id).Find(&addresses)
		ids := make([]int, len(addresses))
		for _, tmp := range addresses {
			ids = append(ids, tmp.Id)
		}
		config.DBHelper.Table("shop_addresses").Where("id IN (?)", ids).Updates(map[string]interface{}{"is_default": 0})
	}

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

	if address.IsDefault {
		u := config.ShowUser(c)
		addresses := []model.ShopAddresses{}
		config.DBHelper.Where("user_id = ?", u.Id).Find(&addresses)
		ids := make([]int, len(addresses))
		for _, tmp := range addresses {
			ids = append(ids, tmp.Id)
		}
		config.DBHelper.Table("shop_addresses").Where("id IN (?)", ids).Updates(map[string]interface{}{"is_default": 0})
	}

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

// 删除地址id
func DeleteAddressById(c *gin.Context)  {
	id := c.PostForm("id")
	address := model.ShopAddresses{}
	addId, _ := strconv.Atoi(id)
	address.Id = addId
	config.DBHelper.Delete(&address)
	c.JSON(200, config.CodeJSON(200, "成功", nil))
}