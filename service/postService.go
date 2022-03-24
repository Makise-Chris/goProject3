package service

import (
	"goProject3/models"
	"goProject3/repository"
)

type PostService interface {
	GetAllPosts(limit, offset int, sort string) ([]models.Post, error)
	GetAllPostsByUserId(userId, limit, offset int, sort string) ([]models.Post, error)
	GetPostByUserId(userId, postId int) (models.Post, error)
	CheckUser(userId int) error
	CreatePost(post models.Post) error
	CheckUserPost(userId, postId int) error
	UpdatePost(post models.Post) error
	DeletePost(post models.Post) error
	ValidatePost(post models.Post) string
}

type PostServiceImpl struct {
	PostRepo repository.PostRepo
}

func NewPostService(r repository.PostRepo) PostService {
	return &PostServiceImpl{
		PostRepo: r,
	}
}

func (p *PostServiceImpl) GetAllPosts(limit, offset int, sort string) ([]models.Post, error) {
	return p.PostRepo.GetAllPosts(limit, offset, sort)
}

func (p *PostServiceImpl) GetAllPostsByUserId(userId, limit, offset int, sort string) ([]models.Post, error) {
	return p.PostRepo.GetAllPostsByUserId(userId, limit, offset, sort)
}

func (p *PostServiceImpl) GetPostByUserId(userId, postId int) (models.Post, error) {
	return p.PostRepo.GetPostByUserId(userId, postId)
}

func (p *PostServiceImpl) CheckUser(userId int) error {
	return p.PostRepo.CheckUser(userId)
}

func (p *PostServiceImpl) CreatePost(post models.Post) error {
	return p.PostRepo.CreatePost(post)
}

func (p *PostServiceImpl) UpdatePost(post models.Post) error {
	return p.PostRepo.UpdatePost(post)
}

func (p *PostServiceImpl) DeletePost(post models.Post) error {
	return p.PostRepo.DeletePost(post)
}

func (p *PostServiceImpl) CheckUserPost(userId, postId int) error {
	return p.PostRepo.CheckUserPost(userId, postId)
}

func (p *PostServiceImpl) ValidatePost(post models.Post) string {
	err := models.Validate.Struct(post)
	if err != nil {
		return "Thiếu caption"
	}
	return ""
}
