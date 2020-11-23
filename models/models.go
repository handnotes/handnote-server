package models

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/handnotes/handnote-server/pkg/setting"
)

var DB *gorm.DB
var logLevel = logger.Error

func InitDatabase() {
	if gin.Mode() == gin.TestMode {
		return
	}

	var err error

	if gin.Mode() == gin.DebugMode || os.Getenv("DEBUG_SQL") == "true" {
		logLevel = logger.Info
	}

	switch setting.Database.Dialect {
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
			setting.Database.Host, setting.Database.Port, setting.Database.User,
			setting.Database.Dbname, setting.Database.Sslmode, setting.Database.Password)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		dsn := setting.Database.SqliteFile
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		log.Fatalf("Invalid database dialect: %s", setting.Database.Dialect)
	}

	if err != nil {
		log.Fatalln(err)
	}

	DB.Logger = logger.Default.LogMode(logLevel)

	err = DB.AutoMigrate(&User{})
	err = DB.AutoMigrate(&Memo{})
	err = DB.AutoMigrate(&Version{})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
