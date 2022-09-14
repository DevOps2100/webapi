package dao

import (
	"go.uber.org/zap"
	"webapi/api/forms"
	"webapi/api/models"
	"webapi/global"
)

func CheckExistHost(host string) (bool, string) {
	var info []models.Host
	result := global.DB.Where("name = ?", host).Find(&info)
	result.Row().Scan(&info)
	//fmt.Printf("信息: %v  长度： %d\n", info, len(info))
	if len(info) >= 1 {
		return true, "数据信息已存在"
	}
	return false, "数据信息不存在"
}

// add host
func AddHost(hostinfo models.Host) (bool, string) {
	NotExist, response := CheckExistHost(hostinfo.Name)
	if NotExist {
		return false, response
	} else {
		result := global.DB.Create(&hostinfo)
		if result.Error != nil {
			zap.L().Info("host add faild")
			return false, "添加失败"
		}
		return true, "添加成功"
	}
}

// 删除主机
func DelHost(hostname string) (bool, string) {
	var host models.Host
	NotExist, response := CheckExistHost(hostname)
	if NotExist {
		return false, response
	}
	host.Name = hostname
	result := global.DB.Where("name = ?", hostname).Delete(&host)
	if result.RowsAffected > 1 {
		return false, "delete failed"
	}
	return true, "success"
}

// 更新主机
func UpdateHost(info forms.Host) (bool, string) {
	var hostinfo models.Host
	NotExist, response := CheckExistHost(info.Name)
	if NotExist {
		return false, response
	} else {
		global.DB.First(&hostinfo)
		hostinfo.Name = info.Name
		hostinfo.Ipaddress = info.Ipaddress
		hostinfo.Password = info.Password
		hostinfo.Memory = info.Memory
		global.DB.Save(&hostinfo)
		return true, "更新成功"
	}
}

// 获取主机(单个)
func GetHost(name string) (*models.Host, string) {
	var info models.Host
	NotExist, response := CheckExistHost(info.Name)
	if NotExist {
		return nil, response
	} else {
		global.DB.Where("name = ?", name).First(&info)
		return &info, ""
	}
}

// 获取主机(全部)
func GetHostAll() []models.Host {
	var host []models.Host
	result := global.DB.Find(&host)
	if result.Error != nil {
		return nil
	}
	if result.RowsAffected < 1 {
		return nil
	}
	return host
}
