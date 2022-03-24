package service

import (
	"goProject3/models"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) ValidateUser(user models.User2) string {
	ret := m.Called(user)
	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}
	return r0
}

func (m *MockUserService) ValidateAuth(auth models.Authentication) string {
	ret := m.Called(auth)
	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}
	return r0
}

func (m *MockUserService) CreateUser(user models.User2) error {
	ret := m.Called(user)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}

func (m *MockUserService) DeleteUser(user models.User2) error {
	ret := m.Called(user)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}

func (m *MockUserService) GetUserById(userId int) (models.User2, error) {
	ret := m.Called(userId)

	var r0 models.User2
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.User2)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockUserService) GetUserByEmail(email string) (models.User2, error) {
	ret := m.Called(email)

	var r0 models.User2
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.User2)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
