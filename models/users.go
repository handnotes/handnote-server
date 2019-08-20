package models

import (
	"fmt"
	"time"
)

// TableName 指定用户表表名.
func (User) TableName() string {
	return "users"
}

// TableName 指定用户表表名.
func (UserForm) TableName() string {
	return "users"
}

// User 定义用户表对应的结构.
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	UserName  string    `json:"user_name" gorm:"size:50;not null;default:''"`
	Phone     string    `json:"phone" gorm:"type:char(11);not null;default:''"`
	Password  string    `json:"password" gorm:"type:char(60);not null;default:''"`
	Email     string    `json:"email" gorm:"size:50;not null;default:''"`
	Address   string    `json:"address" gorm:"size:200;not null;default:''"`
	Gender    int8      `json:"gender" gorm:"not null;default:1"`
	Birth     time.Time `json:"birth" gorm:"not null;type:date;default:'1970-01-01'"`
	AvatarURL string    `json:"avatar_url" gorm:"size:200;not null;default:''"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// UserForm 用户注册表单.
type UserForm struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Address   string    `json:"address"`
	Gender    int8      `json:"gender" binding:"required"`
	Birth     time.Time `json:"birth"`
	AvatarURL string    `json:"avatar_url"`
	Code      int       `json:"code"`
}

// UserEmail 用于接受、验证用户邮件的结构.
type UserEmail struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"user_name"`
}

// GetUserList 获取用户列表.
func GetUserList() (users []User, err error) {
	if err = dbConn.Find(&users).Error; err != nil {
		return
	}
	return
}

// CreateUser is a func to create user.
func CreateUser(user *UserForm) error {
	if err := dbConn.Omit("id", "code").Create(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}
