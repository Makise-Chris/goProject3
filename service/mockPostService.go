package service

import (
	"goProject3/models"

	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) GetAllPosts(limit, offset int, sort string) ([]models.Post, error) {
	ret := m.Called(limit, offset, sort)
	var r0 []models.Post
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]models.Post)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockPostService) GetAllPostsByUserId(userId, limit, offset int, sort string) ([]models.Post, error) {
	ret := m.Called(userId, limit, offset, sort)
	var r0 []models.Post
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]models.Post)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockPostService) GetPostByUserId(userId, postId int) (models.Post, error) {
	ret := m.Called(userId, postId)
	var r0 models.Post
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.Post)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockPostService) CheckUser(userId int) error {
	ret := m.Called(userId)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockPostService) CreatePost(post models.Post) error {
	ret := m.Called(post)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockPostService) UpdatePost(post models.Post) error {
	ret := m.Called(post)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockPostService) DeletePost(post models.Post) error {
	ret := m.Called(post)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockPostService) CheckUserPost(userId, postId int) error {
	ret := m.Called(userId, postId)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockPostService) ValidatePost(post models.Post) string {
	ret := m.Called(post)
	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	return r0
}
