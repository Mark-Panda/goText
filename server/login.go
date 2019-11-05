package server

import (
	"fmt"
	"ginProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	name := c.PostForm("name")
	fmt.Println("入参",name)
	info,err := model.FindMany(name)
	if err != nil {
		fmt.Println("查询错误",err)
	}
	c.JSON(200, gin.H{
		"context": info,
	})
}


func Posting(c *gin.Context)  {
	name := c.PostForm("name")
	passowrd := c.PostForm("password")

	c.JSON(200, gin.H{
		"status": true,
		"name": name,
		"password": passowrd,
	})
}

func MixReq(c *gin.Context)  {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

func Upload(c *gin.Context)  {
	//file  上传的文件文件头
	file, _ := c.FormFile("file")
	//获取文件名
	fileName := file.Filename
	fmt.Println("文件名：", fileName)

	//上传文件到指定路径
	err := c.SaveUploadedFile(file, "static/uploadFile/"+fileName)
	if err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

