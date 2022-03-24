package repository

import (
	"fmt"
	"goProject3/models"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetUserById(int) (models.User2, error)
	GetUserByEmail(string) (models.User2, error)
	CreateUser(models.User2) error
	DeleteUser(models.User2) error
	Migrate() error
}

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (u *UserRepoImpl) Migrate() error {
	fmt.Print("UserRepository...Migrate")
	result := u.Db.AutoMigrate(&models.User2{})
	return result.Error
}

func (u *UserRepoImpl) DeleteUser(user models.User2) error {
	result := u.Db.Unscoped().Delete(&user)
	return result.Error
}

func (u *UserRepoImpl) GetUserById(userId int) (models.User2, error) {
	var dbuser models.User2
	result := u.Db.Where("id = ?", userId).First(&dbuser)
	return dbuser, result.Error
}

func (u *UserRepoImpl) GetUserByEmail(email string) (models.User2, error) {
	var dbuser models.User2
	result := u.Db.Where("email = ?", email).First(&dbuser)
	return dbuser, result.Error
}

func (u *UserRepoImpl) CreateUser(user models.User2) error {
	result := u.Db.Create(&user)
	return result.Error
}
