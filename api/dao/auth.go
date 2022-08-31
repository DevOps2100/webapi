package dao

import (
	"webapi/api/models"
	"webapi/global"
)

// 通过用户名查询用户信息
func GetUserByUsername(username string) models.User {

	// err = DB.Table("peoples").Select("password").Where("username=?", user.Username).First(pwd).Error
	// if err == sql.ErrNoRows{
	// 	fmt.Println(err)
	// 	return errors.New("用户不存在")
	// }
	// if err != nil{
	// 	//查询数据库失败
	// 	return err
	// }
	// //判断密码是否正确
	// if user.Password != pwd.Password{
	// 	err = errors.New("密码错误")
	// 	return err
	// }

	user := models.User{}
	global.DB.Find(&user, global.DB.Where("username = ?", username))
	return user
}
