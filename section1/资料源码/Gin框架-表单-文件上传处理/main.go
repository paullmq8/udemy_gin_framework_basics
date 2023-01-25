package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//32Mib
	// router.MaxMultipartMemory = 8

	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传错误: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("保存错误: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("文件上传成功 %s %s filename=%s", name, email, file.Filename))
	})

	router.Run(":8080")

}
