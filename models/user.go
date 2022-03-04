package models

import "github.com/jinzhu/gorm"

type User2 struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `gorm:"unique" json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
	Posts    []Post `json:"posts"`
}

type User2SignUp struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
