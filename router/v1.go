package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ginProject/server"
)

func NewInit() *gin.Engine {
	e := gin.New()
	v1 := e.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			name := c.Query("name")
			fmt.Println("111111",name)
			c.JSON(200, gin.H{
				"message": name,
			})
		})
		//POST 方法
		v1.POST("/somePost", server.Posting)
		/*
			GET POST 混和使用
			POST /post?id=1234&page=1 HTTP/1.1
			Content-Type: application/x-www-form-urlencoded
		*/
		v1.POST("/getMixPost", server.MixReq)
		//表单上传
		v1.POST("/upload", server.Upload)

		//测试数据库连接
		v1.POST("/database", server.Login)
	}
	return e
}


