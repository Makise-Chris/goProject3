package models

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"

	"goProject3/config"
)

var (
	lock                = &sync.Mutex{}
	Secretkey    string = "secretkeyjwt"
	Connection   *gorm.DB
	Validate     *validator.Validate
	DefaultLimit string = "2"
	DefaultPage  string = "1"
	DefaultSort  string = "asc"
)

//Singleton Pattern
func GetSingleInstance() *gorm.DB {
	if Connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if Connection == nil {
			fmt.Println("Creating single instance database now.")
			Connection = &gorm.DB{}
			Connection = config.GetDatabase()
		} else {
			fmt.Println("Single instance database already created.")
		}
	} else {
		fmt.Println("Single instance database already created.")
	}

	return Connection
}
