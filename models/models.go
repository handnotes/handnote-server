package models

import (
	"fmt"
	"log"

	"github.com/handnotes/handnote-server/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

var DB *gorm.DB

// init 初始化数据库连接
func init() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		setting.Database.Host, setting.Database.Port, setting.Database.User,
		setting.Database.Dbname, setting.Database.Sslmode, setting.Database.Password)
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Memo{})
	DB.AutoMigrate(&Version{})
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
