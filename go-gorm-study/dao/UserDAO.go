package dao

import (
	"go-gorm-study/config"
	"go-gorm-study/model"
	"time"
)

type UserDAO struct {
}

/*func init() {
	if (!config.DB.Migrator().HasTable(&model.User{})) {
		err := config.DB.AutoMigrate(&model.User{})
		if err != nil {
			return
		}
	}
}*/

// 根据主键查询 /*

func (userDAO UserDAO) GetById(id int64) *model.User {
	var user model.User
	config.DB.Find(&user, id)
	return &user
}

// SaveOrUpdate 根据主键 ID /*
func (userDAO UserDAO) SaveOrUpdate(user *model.User) uint {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Deleted = 0
	config.DB.Save(user)
	return user.ID
}
