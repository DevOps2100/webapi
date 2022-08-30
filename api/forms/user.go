package forms

/*
注意几个点:
	1. Captcha的最大值和最小值都为5,应为后续设置获取验证码的位数是5
	2. PasswordLoginForm修改后,login的接口就必须传入验证码和验证id,否则报错
*/

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	// 密码  binding:"required"为必填字段,长度大于3小于20
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

type UserAddForm struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type UserListAllForm struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
}

type UsernameForm struct {
	Username string `form:"username" json:"username" binding:"required"`
}
