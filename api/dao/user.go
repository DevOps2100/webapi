package dao

import (
	"fmt"
	"webapi/api/forms"
	"webapi/api/models"
	"webapi/global"

	"go.uber.org/zap"
)

// 检查用户是否存在
func UserCheckIsExistUserName(username string) (bool, string) {
	var user []models.User
	result := global.DB.Where("username = ?", username).Find(&user)
	result.Row().Scan(&user)
	if len(user) < 1 {
		return true, "用户不存在"
	}
	fmt.Println(user, len(user))
	return false, "用户已存在"
}

// 用户添加
func AddUser(User *forms.UserAddForm) string {
	var user models.User
	user.Username = User.Username
	user.Password = User.Password
	user.Role = User.Role
	user.PhoneNumber = User.PhoneNumber
	user.Avatar = User.AvaTar

	// 先查询用户是否存在，存在则返回存在
	ok, response := UserCheckIsExistUserName(user.Username)
	if !ok {
		zap.L().Info(response)
		return response
	} else {
		result := global.DB.Create(&user)
		if result.Error != nil {
			zap.L().Error("用户添加失败")
			return "添加失败"
		}
		zap.L().Info("用户添加成功")
		return "添加成功"
	}
}

// 用户获取(单个)
func GetUser(Username string) (*models.User, string) {
	var user models.User
	NotExist, response := UserCheckIsExistUserName(Username)
	if NotExist {
		zap.L().Info("查询的用户不存在")
		return &user, response
	} else {
		rows := global.DB.Where("username = ?", Username).Find(&user)
		if rows.RowsAffected < 1 {
			zap.L().Info("查询的用户不存在")
			return &user, response
		} else {
			return &user, "查询成功"
		}

	}
}

// 用户获取(多个)
func GetUserAll() []models.User {
	var user []models.User
	response := global.DB.Find(&user)
	if response.RowsAffected > 0 {
		zap.L().Info("查询成功")
		return user
	}
	return user
}

// 删除用户
func DeleteUser(username string) (bool, string) {
	var user models.User
	Notexist, response := UserCheckIsExistUserName(username)
	if Notexist {
		zap.L().Info("用户不存在,无法执行删除动作")
		return false, response
	} else {
		response := global.DB.Where("username = ?", username).Delete(&user)
		if response.RowsAffected > 1 {
			zap.L().Info("删除失败")
			return true, "删除失败"
		} else {
			zap.L().Info("删除成功")
			return false, "删除成功"
		}
	}
}

// 结构体要符合表字段
type User struct {
	Username    string
	Password    string
	Role        string
	PhoneNumber string
	Avatar      string
}

// 用户修改
func UpdateUser(user forms.UserAddForm) (bool, string) {
	updataFood := User{
		Username:    user.Username,
		Password:    user.Password,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.AvaTar,
	}

	// 根据条件进行更新
	response := global.DB.Model(&User{}).Where("username = ?", user.Username).Updates(&updataFood)
	fmt.Println(response.RowsAffected)
	zap.L().Info("修改成功")
	return true, "修改成功"
}
