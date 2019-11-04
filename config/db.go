package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = "3306"
	DATABASE = "blog"
)
func InitDB()  {
	connInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", connInfo)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		log.Fatal(err)
	}
	DB.SetConnMaxLifetime(100 * time.Second)  //最大连接周期， 超时就CLose
	DB.SetMaxOpenConns(100)  //设置最大连接数
}

