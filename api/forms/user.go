package forms

/*
注意几个点:
	1. Captcha的最大值和最小值都为5,应为后续设置获取验证码的位数是5
	2. PasswordLoginForm修改后,login的接口就必须传入验证码和验证id,否则报错
*/

type UserAddForm struct {
	Username    string `form:"username" json:"username"`         // 用户
	Password    string `form:"password" json:"password"`         // 密码
	Role        string `form:"role" json:"role"`                 // 角色
	PhoneNumber string `form:"phone_number" json:"phone_number"` // 手机号
	AvaTar      string `form:"vavtar" json:"vavtar"`             // 头像
}

type UserListAllForm struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
}

type UsernameForm struct {
	Username string `form:"username" json:"username" binding:"required"`
}

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
