package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var (
	Secretkey    string = "secretkeyjwt"
	Connection   *gorm.DB
	Validate     *validator.Validate
	DefaultLimit string = "2"
	DefaultPage  string = "1"
	DefaultSort  string = "asc"
)
