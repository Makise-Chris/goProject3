package repository

import (
	"goProject3/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Migrate() error {
	ret := m.Called()
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}

func (m *MockUserRepo) CreateUser(user models.User2) error {
	ret := m.Called(user)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}

func (m *MockUserRepo) DeleteUser(user models.User2) error {
	ret := m.Called(user)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}

func (m *MockUserRepo) GetUserById(userId int) (models.User2, error) {
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

func (m *MockUserRepo) GetUserByEmail(email string) (models.User2, error) {
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
