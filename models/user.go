package models

import "time"

type User2 struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `gorm:"unique" json:"email"  validate:"required,email"`
	Password  string     `json:"password" validate:"required"`
	Role      string     `json:"role" validate:"required,oneof=admin user"`
	Posts     []Post     `json:"posts"`
}

type User2SignUp struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
