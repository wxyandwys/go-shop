package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"empolder/src/controller"
)

func main()  {
	router := gin.Default()
	v1 := router.Group("/v1/home")
	{
		v1.GET("", controller.GetHome)
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