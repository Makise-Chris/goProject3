package service

import (
	"goProject3/models"
	"goProject3/repository"
)

type CommentService interface {
	GetAllCommentsByPostId(postId, limit, offset int, sort string) ([]models.Comment, error)
	GetAllCommentsByPostIdAndUserId(postId, userId, limit, offset int, sort string) ([]models.Comment, error)
	CreateComment(comment models.Comment) error
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteComment(comment models.Comment) error
	CheckUserComment(userId, commentId int) error
}

type CommentServiceImpl struct {
	CommentRepo repository.CommentRepo
}

func NewCommentService(r repository.CommentRepo) CommentService {
	return &CommentServiceImpl{
		CommentRepo: r,
	}
}

func (c *CommentServiceImpl) GetAllCommentsByPostId(postId, limit, offset int, sort string) ([]models.Comment, error) {
	return c.CommentRepo.GetAllCommentsByPostId(postId, limit, offset, sort)
}

func (c *CommentServiceImpl) GetAllCommentsByPostIdAndUserId(postId, userId, limit, offset int, sort string) ([]models.Comment, error) {
	return c.CommentRepo.GetAllCommentsByPostIdAndUserId(postId, userId, limit, offset, sort)
}

func (c *CommentServiceImpl) CreateComment(comment models.Comment) error {
	return c.CommentRepo.CreateComment(comment)
}

func (c *CommentServiceImpl) UpdateComment(comment models.Comment) (models.Comment, error) {
	return c.CommentRepo.UpdateComment(comment)
}

func (c *CommentServiceImpl) DeleteComment(comment models.Comment) error {
	return c.CommentRepo.DeleteComment(comment)
}

func (c *CommentServiceImpl) CheckUserComment(userId, commentId int) error {
	return c.CommentRepo.CheckUserComment(userId, commentId)
}
