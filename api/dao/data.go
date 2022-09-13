package dao

import (
	"fmt"
	"webapi/api/forms"
	"webapi/api/models"
	"webapi/global"
)

func CheckExist(datainfo string) (bool, string) {
	var info []models.Data
	result := global.DB.Where("name = ?", datainfo).Find(&info)
	result.Row().Scan(&info)
	//fmt.Printf("信息: %v  长度： %d\n", info, len(info))
	if len(info) >= 1 {
		return true, "数据信息已存在"
	}
	return false, "数据信息不存在"
}

// 数据库信息添加
func DataAdd(INFO forms.Data) (bool, string) {
	//var data models.Data
	var info models.Data
	info.Name = INFO.Name
	info.ConnectSite = INFO.ConnectSite
	info.Port = INFO.Port
	info.Username = INFO.Username
	info.Passwd = INFO.Passwd
	info.DatabaseType = INFO.DatabaseType
	if ok, response := CheckExist(INFO.Name); ok {
		// 开始添加数据信息
		result := global.DB.Create(&info)
		if result.Error != nil {
			fmt.Println("错误信息： ", result.Error)
			return false, response
		}
		return true, "添加成功"

	} else {
		return false, response
	}
}

//  数据库信息删除
func DataDel(name string) (bool, string) {
	var info models.Data
	ok, response := CheckExist(name)
	if !ok {
		return false, response
	}
	result := global.DB.Where("name = ?", name).Delete(&info)
	if result.RowsAffected > 1 {
		return false, "删除失败"
	}
	return true, "删除成功"
}

//   数据库信息更新
func DataUpdate(info forms.Data) (bool, string) {
	var data models.Data
	ok, response := CheckExist(info.Name)
	if !ok {
		return false, response
	}
	result := global.DB.Model(&data).Where("name = ?", info.Name).Updates(&info)
	fmt.Println(result.RowsAffected)
	return true, "修改成功"
}

//  数据库信息获取
func DataGet(sd forms.Data) (*models.Data, string) {
	var rd models.Data
	ok, response := CheckExist(sd.Name)
	if !ok {
		return nil, response
	}
	rows := global.DB.Where("name = ?", sd.Name).Find(&rd)
	if rows.RowsAffected < 1 {
		return nil, response
	}
	return &rd, "获取成功"
}
