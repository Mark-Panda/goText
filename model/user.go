package model

import (
	"ginProject/config"
	"fmt"
	"ginProject/schema"
)
func FindOne(name string) (users *schema.User, err error){
	user := new(schema.User)
	sql := `select * from user where user_name = ?`
	row := config.DB.Exec(sql, name)
	fmt.Println("具体",row)
	fmt.Println("数据信息", user)
	return
}

func FindMany(name string) (users []*schema.User, err error) {
	err = config.DB.Where("username=?",name).Find(&users).Error
	//fmt.Println("数据信息", row)
	return
}
