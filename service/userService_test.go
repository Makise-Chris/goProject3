package service

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"goProject3/models"
	"goProject3/repository"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUser := models.User2{
		Email:    "bob@bob.com",
		Name:     "Bobby Bobson",
		Password: "Nam12345",
		Role:     "user",
	}

	mockUserRepo := new(repository.MockUserRepo)
	userService := NewUserService(mockUserRepo)

	mockUserRepo.On("CreateUser", mockUser).Return(nil)
	error := userService.CreateUser(mockUser)

	assert.Equal(t, nil, error)
	mockUserRepo.AssertExpectations(t)
}

//TestValidateUser lỗi
func TestValidateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUser := models.User2{
		Email:    "bob@bob.com",
		Name:     "Bobby Bobson",
		Password: "Nam12345",
		Role:     "user",
	}

	mockUserRepo := new(repository.MockUserRepo)
	userService := NewUserService(mockUserRepo)

	message := userService.ValidateUser(mockUser)

	if message != "" {
		t.Errorf("Want " + "" + ", got " + message)
	}
	mockUserRepo.AssertExpectations(t)
}
