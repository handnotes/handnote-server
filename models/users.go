package models

import (
	"time"
)

// TableName 指定用户表表名.
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构.
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	UserName  string    `json:"user_name" gorm:"size:50;not null;default:''"`
	Phone     string    `json:"phone" gorm:"type:char(11);not null;default:''"`
	Email     string    `json:"email" gorm:"size:50;not null;default:''"`
	Address   string    `json:"address" gorm:"size:200;not null;default:''"`
	Gender    int8      `json:"gender" gorm:"not null;default:1"`
	AvatarURL string    `json:"avatar_url" gorm:"size:200;not null;default:''"`
	Status    int8      `json:"status" gorm:"not null;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// GetUserList 获取用户列表.
func GetUserList() (users []User, err error) {
	if err = dbConn.Find(&users).Error; err != nil {
		return
	}
	return
}
