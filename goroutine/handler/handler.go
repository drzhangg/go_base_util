package handler

import (
	"errors"
	"go_base_util/goroutine/model"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) AddUser() error {
	model.Db.Table("user").Where("name = ?", u.Name).First(&u)
	if u.Id > 0 {
		return errors.New("该用户已注册")
	}
	return model.Db.Table("user").Create(&u).Error
}
