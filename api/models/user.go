package models

import (
	"time"
)

// 用户表字段
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Password  string    `json:"password" gorm:"password"`
	Username  string    `json:"username" gorm:"username"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

//用户表名称
func (User) TableName() string {
	return "users"
}
