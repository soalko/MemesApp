package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host='188.120.255.223' user='memadmin' password='memadmin229339!?!' dbname='mem_db' port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Db", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db
}
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Meme{},
		&User{},
		&UserRole{},
		&Tag{},
		&Category{},
		&TagMeme{},
		&CategoryMeme{},
	)
}
