package models

import (
	"fmt"
	"time"

	"github.com/handnotes/handnote-server/pkg/util"
	"github.com/jinzhu/gorm"
)

// TableName 指定用户表表名
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构
type User struct {
	ID        uint      `form:"id" json:"id" gorm:"primary_key;not null;auto_increment"`
	Phone     string    `form:"phone" json:"phone" gorm:"type:char(11);unique;default:''"`
	Email     string    `form:"email" json:"email" gorm:"size:50;unique_index;not null;default:''"`
	UserName  string    `form:"user_name" json:"user_name" gorm:"size:50;not null;default:''"`
	Password  string    `form:"password" json:"password" gorm:"type:char(60);not null;default:''"`
	Gender    int8      `form:"gender" json:"gender" gorm:"not null;default:0"`
	Birth     time.Time `form:"birth" json:"birth" gorm:"not null;type:date;default:'1990-01-01'"`
	AvatarURL string    `form:"avatar_url" json:"avatar_url" gorm:"size:200;default:''"`
	CreatedAt time.Time `form:"created_at" json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// GetUserByEmail 通过邮箱获取单个用户
func GetUserByEmail(email string) (user User, err error) {
	if err = DB.Where(User{Email: email}).Find(&user).Error; err != nil {
		return
	}
	return
}

// BeforeSave ...
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if user.Password, err = util.GeneratePassword(user.Password); err != nil {
		return
	}
	_ = scope.SetColumn("Password", user.Password)
	return
}

// SaveUser 保存用户信息，包括创建/更新
func SaveUser(user *User) error {
	if err := DB.Save(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}
