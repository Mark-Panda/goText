package model

import (
	"ginProject/config/"
	"fmt"
	"ginProject/schema"
)
func FindOne(name string) {
	user := new(schema.User)
	sql := `select * from user where user_name = ?`
	row, err := config.DB.Query(sql,name)
	if err != nil {
		fmt.Println("查找失败")
	}
	fmt.Println("具体",row)
	fmt.Println("数据信息", user)
}
