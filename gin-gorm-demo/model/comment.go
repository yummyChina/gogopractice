package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null;type:text" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}
