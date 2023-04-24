package main

import (
	"fmt"
	"go-gorm-study/dao"
	"go-gorm-study/model"
)

func main() {
	user := dao.UserDAO{}.GetById(1)
	fmt.Println("user=", *user)
	user2 := &model.User{
		Name:      "test002",
		Phone:     "123456789",
		LoginName: "1234567",
		PassWord:  "123456",
		Address:   "北京-北京-朝阳",
	}
	user2.ID = 1
	id := dao.UserDAO{}.SaveOrUpdate(user2)
	fmt.Println("插入成功=", id)
}
