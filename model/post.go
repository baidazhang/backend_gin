package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserId     uint `json:"user_id" gorm:"not null"`
	CategoryId uint `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
}
