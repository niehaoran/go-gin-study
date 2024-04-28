package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由
	router := gin.Default()
	//设置路由
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	//post
	router.POST("/xxxpost", func(c *gin.Context) {
		c.String(http.StatusOK, "Received a POST request")
	})
	router.PUT("/xxxput")
	//设置监听端口
	router.Run(":8080")
}
