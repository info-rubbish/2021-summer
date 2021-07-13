package database

import (
	"bytes"
	"main/src/config"
	"main/src/tokens"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Created    time.Time `gorm:"autoCreateTime"`
	Permission uint      `gorm:"not null"`
	ID         string    `gorm:"primaryKey;unique;not null"`
	Name       string    `gorm:"primaryKey;unique;not null"`
	Password   []byte    `gorm:"not null"`
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
func UserLogin(name string, password string) (*User, error) {
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
	var hashPassword []byte
	if c.Password != "" {
		hashPassword = tokens.Hash([]byte(c.Password))
	}
	if err := DB.Model(&User{}).Where("id=?", id).Updates(&User{
		Name:     c.Name,
		Password: hashPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(name, password string) error {
	if err := DB.Create(&User{
		Permission: 1,
		Name:       name,
		Password:   tokens.Hash([]byte(password)),
	}).Error; err != nil {
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
