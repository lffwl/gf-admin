package model

type AuthLoginReq struct {
	AuthLoginInput
	UserName string `v:"required"`           // 用户名
	Password string `v:"required|password2"` // 密码
}

type AuthLoginInput struct {
	UserName string
	Password string
}
