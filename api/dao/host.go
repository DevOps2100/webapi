package dao

import (
	"fmt"
	"go.uber.org/zap"
	"webapi/api/forms"
	"webapi/api/models"
	"webapi/global"
)

// 涓绘満娣诲姞
func AddHost(hostinfo forms.Host) (bool, error) {
	result := global.DB.Create(&hostinfo)
	if result.Error != nil {
		zap.L().Info("host add faild")
		return false, result.Error
	}
	return true, nil
}

// 涓绘満鍒犻櫎
func DelHost(hostname string) (bool, string) {
	var host models.Host
	host.Name = hostname
	response := global.DB.Where("name = ?", hostname).Delete(&host)
	if response.RowsAffected > 1 {
		return true, "鍒犻櫎澶辫触"
	}
	return true, "鍒犻櫎澶辫触"
}

// 涓绘満鏇存柊
func UpdateHost(info forms.Host) {
	var hostinfo models.Host
	global.DB.First(&hostinfo)
	hostinfo.Name = info.Name
	hostinfo.Ipaddress = info.Ipaddress
	hostinfo.Password = info.Password
	hostinfo.Memory = info.Memory
	global.DB.Save(&hostinfo)
	return
}

// 涓绘満鑾峰彇(鍗曚釜)
func GetHost(name string) *models.Host {
	var info models.Host
	global.DB.Where("name = ?", name).First(&info)
	return &info
}

// 涓绘満鑾峰彇(鎵�鏈�)
func GetHostAll() []models.Host {
	var host []models.Host
	result := global.DB.Find(&host)
	if result.Error != nil {
		fmt.Println("璇锋眰閿欒")
		return nil
	}
	if result.RowsAffected < 1 {
		fmt.Println("鏁版嵁涓虹┖")
		return nil
	}
	return host
}
