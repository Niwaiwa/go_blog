package db

import (
	"go_blog/common"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBm *gorm.DB

func InitDBm() {
	configset := common.GetConfig()
	dbHost := configset.DB.Host
	dbPort := configset.DB.Port
	dbUsername := configset.DB.Username
	dbPassword := configset.DB.Password
	dbName := configset.DB.Name

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
		// DefaultStringSize: 256, // default size for string fields
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		// SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	dbpool, err := database.DB()
	if err != nil {
		panic(err)
	}
	dbpool.SetConnMaxLifetime(time.Duration(3600) * time.Second) // 每條連線的存活時間
	dbpool.SetMaxOpenConns(100)                                  // 最大連線數
	dbpool.SetMaxIdleConns(10)
	DBm = database
}
