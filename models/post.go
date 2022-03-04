package models

import "time"

type Post struct {
	//gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Caption   string     `json:"caption" validate:"required"`
	Image     string     `json:"image"`
	User2ID   int        `json:"userid"`
}

type PostSwagger struct {
	Caption string `json:"caption"`
	Image   string `json:"image"`
}
