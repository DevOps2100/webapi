package models

import (
	"time"
)

// 用户表字段
type Data struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"name"`
	DatabaseType string    `json:"database_type" gorm:"database_type"`
	ConnectSite  string    `json:"connect_site" gorm:"connect_site"`
	Username     string    `json:"username" gorm:"username"`
	Passwd       string    `json:"passwd" gorm:"passwd"`
	Port         string    `json:"port" gorm:"port"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
}

//用户表名称
func (Data) TableName() string {
	return "database_info"
}
