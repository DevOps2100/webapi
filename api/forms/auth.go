package forms

/*
注意几个点:
	1. Captcha的最大值和最小值都为5,应为后续设置获取验证码的位数是5
	2. PasswordLoginForm修改后,login的接口就必须传入验证码和验证id,否则报错
*/

type LoginForm struct {
	Password string `form:"password" x-www-form-urlencoded:"password"  json:"password"`
	Username string `form:"username" x-www-form-urlencoded:"username"  json:"username" binding:"required"`
}
