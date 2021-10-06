package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

// 查询收藏
func ShowCollection(c *gin.Context)  {
	userToken :=  c.Request.Header.Get("AccessToken")
	conn := config.RedisDefaultPool.Get()
	user, _ := redis.Bytes( conn.Do("get", userToken) )
	u := model.ShopUsers{}
	json.Unmarshal(user, &u)
	collection := model.ShopsCollections{}
	config.DBHelper.Where("user_id = ?", u.Id).Find(&collection)
	data := map[string]interface{}{}
	data["collection"] = collection
	c.JSON(200, config.CodeJSON(200, "成功", data))
	
}