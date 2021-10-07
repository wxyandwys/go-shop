package config

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"encoding/json"
)
// 查询用户id
func ShowUser(c *gin.Context) model.ShopUsers {
	userToken :=  c.Request.Header.Get("AccessToken")
	conn := RedisDefaultPool.Get()
	user, _ := redis.Bytes( conn.Do("get", userToken) )
	u := model.ShopUsers{}
	json.Unmarshal(user, &u)
	return u
}