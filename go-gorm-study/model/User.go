package model

type User struct {
	BaseModel
	Name      string
	Phone     string
	PassWord  string
	LoginName string
	Address   string
}

func (User) TableName() string {
	return "t_test"
}
