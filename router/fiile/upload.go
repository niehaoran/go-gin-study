package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "upload failed") //500是服务器内部错误
		}
		c.SaveUploadedFile(file, file.Filename) //保存文件到服务器
		c.String(http.StatusOK, "upload success")
	})

	router.Run(":8080")
}
