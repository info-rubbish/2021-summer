package database

import (
	"bytes"
	"main/src/config"
	"main/src/tokens"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Created  time.Time `gorm:"autoCreateTime"`
	ID       string    `gorm:"primaryKey"`
	Name     string    `gorm:"primaryKey"`
	Password []byte
}

type UserConfig struct {
	Name     string
	Password string
}

func (s *User) BeforeCreate(tx *gorm.DB) error {
	s.ID = tokens.RandomID(config.RandomIDLength)
	return nil
}

// return user,error
func UserLogin(s *gin.Context, name string, password string) (*User, error) {
	var user User
	if err := DB.First(&user, "name=?", name).Error; err != nil {
		return nil, err
	}
	if !bytes.Equal(tokens.Hash([]byte(password)), user.Password) {
		return nil, ErrPasswordNotMatch
	}
	return &user, nil
}

func ChangeUserInfo(id string, c UserConfig) error {
	if err := DB.Model(&User{}).Where("id=?", id).Updates(&User{
		Name:     c.Name,
		Password: tokens.Hash([]byte(c.Password)),
	}).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(name, password string) error {
	if err := DB.Create(&User{Name: name, Password: tokens.Hash([]byte(password))}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	if err := DB.Where("id=?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

// cache fn
func UserInfoByID(id string) (interface{}, error) {
	var user User
	if err := DB.First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UserInfoByName(name string) (interface{}, error) {
	var user User
	if err := DB.First(&user, "name=?", name).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
