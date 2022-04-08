package models

import "time"

type Comment struct {
	//gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Comment   string     `json:"comment"`
	User2ID   int        `json:"userid"`
	PostID    int        `json:"postid"`
	User      User2      `json:"user"`
}
