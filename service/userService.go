package service

import (
	"goProject3/models"
	"goProject3/repository"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetUserById(int) (models.User2, error)
	GetUserByEmail(string) (models.User2, error)
	CreateUser(models.User2) error
	DeleteUser(models.User2) error
	ValidateUser(models.User2) string
	ValidateAuth(auth models.Authentication) string
}

type UserSerViceIpml struct {
	UserRepo repository.UserRepo
}

func NewUserService(r repository.UserRepo) UserService {
	return &UserSerViceIpml{
		UserRepo: r,
	}
}

func (u *UserSerViceIpml) DeleteUser(user models.User2) error {
	return u.UserRepo.DeleteUser(user)
}

func (u *UserSerViceIpml) GetUserById(userId int) (models.User2, error) {
	return u.UserRepo.GetUserById(userId)
}

func (u *UserSerViceIpml) GetUserByEmail(email string) (models.User2, error) {
	return u.UserRepo.GetUserByEmail(email)
}

func (u *UserSerViceIpml) ValidateUser(user models.User2) string {
	err := models.Validate.Struct(user)
	if err != nil {
		var message string

		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "oneof" {
				message = message + "Nhập sai role (admin hoặc user). "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}

		return message
	}
	return ""

}

func (u *UserSerViceIpml) CreateUser(user models.User2) error {
	return u.UserRepo.CreateUser(user)
}

func (u *UserSerViceIpml) ValidateAuth(auth models.Authentication) string {
	err := models.Validate.Struct(auth)
	if err != nil {
		var message string

		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}
		return message
	}
	return ""
}
