package database

import (
	"log"
	"main/src/tokens"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			Colorful: true,
		}),
	})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.First(nil, "name=?", "paula").Error; err != nil && err == gorm.ErrRecordNotFound {
		user := &User{
			Name:     "paula",
			Password: tokens.Hash([]byte("20040623")),
		}
		if err := db.Create(&user).Error; err != nil {
			panic(err)
		}
	}
	DB = db
}
