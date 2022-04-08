package repository

import (
	"fmt"
	"goProject3/models"

	"github.com/jinzhu/gorm"
)

type CommentRepo interface {
	Migrate() error
	GetAllCommentsByPostId(postId, limit, offset int, sort string) ([]models.Comment, error)
	GetAllCommentsByPostIdAndUserId(postId, userId, limit, offset int, sort string) ([]models.Comment, error)
	CreateComment(comment models.Comment) error
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteComment(comment models.Comment) error
	CheckUserComment(userId, commentId int) error
}

type CommentRepoIpml struct {
	Db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &CommentRepoIpml{
		Db: db,
	}
}

func (c *CommentRepoIpml) Migrate() error {
	fmt.Print("CommentRepository...Migrate")
	result := c.Db.AutoMigrate(&models.Comment{})
	return result.Error
}

func (c *CommentRepoIpml) GetAllCommentsByPostId(postId, limit, offset int, sort string) ([]models.Comment, error) {
	var comments []models.Comment

	queryBuider := c.Db.Limit(limit).Offset(offset).Order("id " + sort)
	result := queryBuider.Model(&models.Comment{}).Where("post_id = ?", postId).Find(&comments)

	return comments, result.Error
}

func (c *CommentRepoIpml) GetAllCommentsByPostIdAndUserId(postId, userId, limit, offset int, sort string) ([]models.Comment, error) {
	var comments []models.Comment

	queryBuider := c.Db.Limit(limit).Offset(offset).Order("id " + sort)
	result := queryBuider.Model(&models.Comment{}).Where("post_id = ? AND user2_id = ?", postId, userId).Find(&comments)

	return comments, result.Error
}

func (c *CommentRepoIpml) GetNewestCommentByPostIdAndUserId(postId, userId, limit, offset int, sort string) (models.Comment, error) {
	var comments models.Comment

	queryBuider := c.Db.Limit(limit).Offset(offset).Order("id desc")
	result := queryBuider.Model(&models.Comment{}).Where("post_id = ? AND user2_id = ?", postId, userId).First(&comments)

	return comments, result.Error
}

func (c *CommentRepoIpml) CreateComment(comment models.Comment) error {
	result := c.Db.Create(&comment)
	return result.Error
}

func (c *CommentRepoIpml) UpdateComment(comment models.Comment) (models.Comment, error) {
	result := c.Db.Save(&comment)
	return comment, result.Error
}

func (c *CommentRepoIpml) DeleteComment(comment models.Comment) error {
	result := c.Db.Unscoped().Delete(&comment)
	return result.Error
}

func (c *CommentRepoIpml) CheckUserComment(userId, commentId int) error {
	var comment models.Comment

	subQuery := c.Db.Model(&models.Comment{}).Where("user2_id = ?", userId)
	result := subQuery.Where("id = ?", commentId).Find(&comment)

	return result.Error
}
