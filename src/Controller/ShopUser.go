package controller

import (
	"github.com/gin-gonic/gin"
	"empolder/src/model"
	"empolder/src/config"
	"github.com/garyburd/redigo/redis"
	"log"
	"encoding/json"
)

func GetUser(c *gin.Context)  {
	conn := config.RedisDefaultPool.Get()
	userToken := c.PostForm("userToken")
	//取出key值的时候使用redis.Bytes()解析之后再用json.Unmarshal反序列化得到结果。
	user, err := redis.Bytes(conn.Do("get", userToken))
	
	if err != nil {
		log.Fatalln(err)
	}
	
	u := model.ShopUsers{}
	json.Unmarshal(user, &u)
	u.Password = ""
	data := map[string]interface{}{}
	data["user"] = u
	c.JSON(200, config.CodeJSON(200, "成功", data))
}