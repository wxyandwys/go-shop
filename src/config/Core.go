package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/garyburd/redigo/redis"
)
/*
func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		}
			c.Next()
	}
}
*/
// 跨域
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
// 普通权限
func MustLogin() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		if status :=  c.Request.Header.Get("Token"); status == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"text": "没有权限1", "msg": status} )
			c.Abort()
		} else {
			if c.Request.Header.Get("Token") != "11111" {
				c.JSON(http.StatusUnauthorized, "没有权限2")
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
// 用户登录后权限
func LoginUser() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		status :=  c.Request.Header.Get("AccessToken")

		conn := RedisDefaultPool.Get()
		is_key_exit, _ := redis.Bool(conn.Do("EXISTS",status))
		if !is_key_exit {
			c.JSON(402, gin.H{"text": "用户超时，请重新登录", "msg": status} )
			c.Abort()
		} else {
			c.Next()
		}
	
	}
}
