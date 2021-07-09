package database

import (
	"bytes"
	"main/src/config"
	"main/src/tokens"
	"net/http"
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
		return nil, &config.HttpErr{
			Code: http.StatusUnauthorized,
		}
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

// cache fn
func UserInfo(id string) (interface{}, error) {
	var user User
	if err := DB.First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
