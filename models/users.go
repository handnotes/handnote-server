package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/handnotes/handnote-server/pkg/util"
)

func (User) TableName() string {
	return "users"
}

type User struct {
	ID        uint      `form:"id" json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Phone     string    `form:"phone" json:"phone" gorm:"type:char(11);default:''"`
	Email     string    `form:"email" json:"email" gorm:"size:50;uniqueIndex;not null;default:''"`
	UserName  string    `form:"user_name" json:"user_name" gorm:"size:50;not null;default:''"`
	Password  string    `form:"password" json:"password" gorm:"type:char(60);not null;default:''"`
	Gender    int8      `form:"gender" json:"gender" gorm:"not null;default:0"`
	Birth     time.Time `form:"birth" json:"birth" gorm:"not null;type:date;default:'1990-01-01'"`
	AvatarURL string    `form:"avatar_url" json:"avatar_url" gorm:"size:200;default:''"`
	CreatedAt time.Time `form:"created_at" json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" gorm:"not null;default:current_timestamp"`
}

func GetUserByEmail(email string) (user User, err error) {
	err = DB.Where(User{Email: email}).First(&user).Error
	return
}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	user.Password = util.GeneratePassword(user.Password)
	err = db.Model(user).Set("password", user.Password).Error
	return
}

// SaveUser 保存用户信息，包括创建/更新
func SaveUser(user *User) (err error) {
	err = DB.Save(user).Error
	return
}
