package main

import (
	"ginProject/router"
	"github.com/gin-gonic/gin"
	"ginProject/config"
)

func main()  {
	engine := gin.New()
	router.NewInit(engine)
	config.InitDB()
}
