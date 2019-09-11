package models

import (
	"fmt"
	"time"
)

// TableName 指定备忘录/便笺表表名
func (Memo) TableName() string {
	return "memos"
}

// Memo 定义备忘录/便笺表对应的结构
type Memo struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	Name      string    `json:"name" gorm:"size:200;not null;default:''"`
	Content   string    `json:"content" gorm:"not null;default:''"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// GetMemoList 获取备忘录/便笺列表
func GetMemoList() (memos []Memo) {
	dbConn.Find(&memos)
	return
}

// SaveMemo 保存备忘录/便笺信息，包括创建/更新
func SaveMemo(memo *Memo) error {
	if err := dbConn.Save(memo).Error; err != nil {
		return err
	}
	fmt.Println(memo)
	return nil
}
