package main

import (
	"github.com/gin-gonic/gin"
	// 框架
	"net/http"
	"log"
	"empolder/src/controller"

	"empolder/src/config"
	// 跨域
	//"github.com/gin-contrib/sessions"
	//"github.com/garyburd/redigo/redis"
)

func main()  {
	// conn := config.RedisDefaultPool.Get()
	router := gin.Default()

	router.Use(config.Session("topgoer"))
	router.GET("/captcha/:value", func(c *gin.Context) {
		//value := c.Param("value")
		//c.JSON(200, gin.H{"value": value})
		config.Captcha(c, 4)
	})
	/*
	router.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if config.CaptchaVerify(c, value) {
				c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
				c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
	*/
	router.Use(config.Cors())
	router.Use(config.MustLogin())
	v1 := router.Group("/v1/home")
	{
		v1.GET("", controller.GetHome)
		v1.GET("/GetShopCategoryList", controller.GetShopCategoryList)
		v1.GET("/GetShopList", controller.GetShopList)
		v1.GET("/GetShop", controller.GetShop)
		v1.GET("/GetShopTree", controller.GetShopTree)
		v1.GET("/GetShopListSku", controller.GetShopListSku)
		v1.GET("/GetShopParameters", controller.GetShopParameters)
		v1.GET("/GetShopCategoryListParent", controller.GetShopCategoryListParent)
		v1.GET("/GetShopCategoryListChildren", controller.GetShopCategoryListChildren)
		v1.GET("/GetShopCategoryListChildrenById", controller.GetShopCategoryListChildrenById)
		// 登录
		v1.POST("/ShopLogin/:img", func (c *gin.Context)  {
			value := c.PostForm("value")
			controller.ShopLogin(c, value)
		})

	}

	server := &http.Server{
		Addr: ":8901",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("服务器启动失败")
	}
}