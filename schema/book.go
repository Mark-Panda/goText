package schema

import "time"

type Book struct {
	Id int `gorm:"primary_key"`
	Name string
	Price float32
	Author string
	CreatedAt time.Time
}

func (User)FineOne()  {

}
