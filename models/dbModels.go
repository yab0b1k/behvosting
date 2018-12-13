package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Num    int    `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"size:255"`
	Pass   string
	Images []Image
}

type Image struct {
	ID         int
	UserID     uint `gorm:"index"`
	ImageID    string
	Comment    string
	CreateTime time.Time
}
