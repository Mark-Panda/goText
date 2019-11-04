package main

import (
	"ginProject/router"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.New()
	router.NewInit(engine)
}
