package database

import (
	"main/src/config"
	"main/src/tokens"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	Created     time.Time `gorm:"autoCreateTime"`
	ID          string    `gorm:"primaryKey;unique;not null"`
	Author      string    `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
}

func (s *Course) BeforeCreate(tx *gorm.DB) error {
	if err := CheckPermission(s.Author); err != nil {
		return err
	}
	s.ID = tokens.RandomID(config.RandomIDLength)
	return nil
}

func (s *Course) BeforeUpdate(tx *gorm.DB) error {
	if err := CheckPermission(s.Author); err != nil {
		return err
	}
	return nil
}

func (s *Course) BeforeDelete(tx *gorm.DB) error {
	if err := CheckPermission(s.Author); err != nil {
		return err
	}
	return nil
}

type CourseConfig struct {
	Author      string
	Title       string
	Description string
	Content     string
}

func CreateCourse(c *CourseConfig) (string, error) {
	course := &Course{
		Author:      c.Author,
		Title:       c.Title,
		Description: c.Description,
		Content:     c.Content,
	}
	if err := DB.Create(course).Error; err != nil {
		return "", err
	}
	DB.Save(course)
	return course.ID, nil
}

func DeleteCourse(id, author string) error {
	course := &Course{}
	if err := DB.First(course, "id=?", id).Error; err != nil {
		return err
	}
	if course.Author != author {
		return ErrPermissionDenied
	}
	if err := DB.Delete(course).Error; err != nil {
		return err
	}
	return nil
}

func GetCourse(id string) (interface{}, error) {
	course := &Course{}
	if err := DB.First(course, "id=?", id).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func ChangeCourse(id, author string, c *CourseConfig) error {
	if err := DB.Model(&Course{}).Where("id=? AND author = ?", id, author).Updates(&Course{
		Title:       c.Title,
		Description: c.Description,
		Content:     c.Content,
	}).Error; err != nil {
		return err
	}
	return nil
}
