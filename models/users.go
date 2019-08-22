package models

import (
	"fmt"
	"time"

	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/jinzhu/gorm"
)

// TableName 指定用户表表名
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	Phone     string    `json:"phone" gorm:"type:char(11);not null;default:''"`
	Email     string    `json:"email" gorm:"size:50;not null;default:''"`
	UserName  string    `json:"user_name" gorm:"size:50;not null;default:''"`
	Password  string    `json:"password" gorm:"type:char(60);not null;default:''"`
	Address   string    `json:"address" gorm:"size:200;not null;default:''"`
	Gender    int8      `json:"gender" gorm:"not null;default:1"`
	Birth     time.Time `json:"birth" gorm:"not null;type:date;default:'1970-01-01'"`
	AvatarURL string    `json:"avatar_url" gorm:"size:200;not null;default:''"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// GetUserByEmail 通过邮箱获取单个用户
func GetUserByEmail(email string) (user User, err error) {
	if err = dbConn.Where(User{Email: email}).Find(&user).Error; err != nil {
		return
	}
	return
}

// BeforeSave ...
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if user.Password, err = util.GeneratePassword(user.Password); err != nil {
		return
	}
	scope.SetColumn("Password", user.Password)
	return
}

// SaveUser 保存用户信息，包括创建/更新
func SaveUser(user *User) error {
	if err := dbConn.Save(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}
