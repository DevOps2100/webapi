package models

import (
	"time"
)

// 用户表字段
type Host struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"name"`
	Ipaddress string    `json:"ipaddress" gorm:"ipaddress"`
	Password  string    `json:"password" gorm:"password"`
	Cpu       string    `json:"cpu" gorm:"cpu"`
	Memory    string    `json:"memory" gorm:"memory"` // 建议使用monion
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

//用户表名称
func (Host) TableName() string {
	return "hosts"
}
