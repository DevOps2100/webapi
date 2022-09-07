package models

import (
	"time"
)

// 用户表字段
type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Password    string    `json:"password" gorm:"password"`
	Username    string    `json:"username" gorm:"username"`
	Role        string    `json:"role" gorm:"role"`
	PhoneNumber string    `json:"phoneNumber" gorm:"phoneNumber"`
	Avatar      string    `json:"avatar" gorm:"avatar"` // 建议使用monion
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

//用户表名称
func (User) TableName() string {
	return "users"
}
