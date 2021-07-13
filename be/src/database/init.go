package database

import (
	"fmt"
	"main/src/tokens"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"))
	db.Logger.LogMode(logger.Info)
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}, &Course{}); err != nil {
		panic(err)
	}
	user := &User{
		Permission: 2,
		Name:       "paula",
		Password:   tokens.Hash([]byte("20040623")),
	}
	if err := db.First(user, "name=?", "paula").Error; err != nil {
		if err := db.Create(user).Error; err != nil {
			panic(err)
		}
	}
	fmt.Printf("Test Account\n\tName: %s\n", user.Name)
	DB = db
}
