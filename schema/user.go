package schema

type User struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Status int `json:"status"`
	CreateTime int64 `json:"create_time"`
}