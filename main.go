package main

import (
	"github.com/gin-gonic/gin"
	// 框架
	"net/http"
	"log"
	"empolder/src/controller"
	
	"empolder/src/config"
	// 跨域
)

func main()  {
	router := gin.Default()
	router.Use(config.Cors())
	v1 := router.Group("/v1/home")
	{
		v1.GET("", controller.GetHome)
		v1.GET("/GetShopCategoryList", controller.GetShopCategoryList)
		v1.GET("/GetShopList", controller.GetShopList)
		v1.GET("/GetShop", controller.GetShop)
		v1.GET("/GetShopTree", controller.GetShopTree)
		v1.GET("/GetShopListSku", controller.GetShopListSku)
		v1.GET("/GetShopParameters", controller.GetShopParameters)
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