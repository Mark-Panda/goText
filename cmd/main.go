package main

import (
	"ginProject/config"
	"ginProject/router"
)

func main()  {

	config.Init()
	r := router.NewInit()
	r.Run()
}
