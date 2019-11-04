package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
const (
	UserName = "root"
	PassWord = ""
	DataBase = "blog"
)

var DB *gorm.DB
func Init()  {
	connInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", UserName, PassWord, DataBase)
	db, err := gorm.Open("mysql",connInfo)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(100 * time.Second)
	defer DB.Close()
}

