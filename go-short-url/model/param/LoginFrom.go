package param

type LoginFrom struct {
	UserName string `form:"UserName" binding:"required"`
	PassWord string `form:"PassWord" `
}
