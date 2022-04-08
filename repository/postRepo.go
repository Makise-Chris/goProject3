package repository

import (
	"fmt"
	"goProject3/models"

	"github.com/jinzhu/gorm"
)

type PostRepo interface {
	GetAllPosts(limit, offset int, sort string) ([]models.Post, error)
	GetPostById(postId int) (models.Post, error)
	GetAllPostsByUserId(userId, limit, offset int, sort string) ([]models.Post, error)
	GetNewestPostByUserId(userId int) (models.Post, error)
	GetPostByUserId(userId, postId int) (models.Post, error)
	CheckUser(userId int) error
	CreatePost(post models.Post) error
	CheckUserPost(userId, postId int) error
	Migrate() error
	UpdatePost(post models.Post) (models.Post, error)
	DeletePost(post models.Post) error
}

type PostRepoIpml struct {
	Db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &PostRepoIpml{
		Db: db,
	}
}

func (p *PostRepoIpml) Migrate() error {
	fmt.Print("PostRepository...Migrate")
	result := p.Db.AutoMigrate(&models.Post{})
	return result.Error
}

func (p *PostRepoIpml) GetAllPosts(limit, offset int, sort string) ([]models.Post, error) {
	var posts []models.Post

	queryBuider := p.Db.Limit(limit).Offset(offset).Order("id " + sort)
	result := queryBuider.Model(&models.Post{}).Find(&posts)

	return posts, result.Error
}

func (p *PostRepoIpml) GetPostById(postId int) (models.Post, error) {
	var dbpost models.Post
	result := p.Db.Where("id = ?", postId).First(&dbpost)
	return dbpost, result.Error
}

func (p *PostRepoIpml) GetAllPostsByUserId(userId, limit, offset int, sort string) ([]models.Post, error) {
	var posts []models.Post

	queryBuider := p.Db.Limit(limit).Offset(offset).Order("id " + sort)
	result := queryBuider.Model(&models.Post{}).Where("user2_id = ?", userId).Find(&posts)

	return posts, result.Error
}

func (p *PostRepoIpml) GetNewestPostByUserId(userId int) (models.Post, error) {
	var post models.Post
	result := p.Db.Model(&models.Post{}).Where("user2_id = ?", userId).Order("id desc").First(&post)

	return post, result.Error
}

func (p *PostRepoIpml) GetPostByUserId(userId, postId int) (models.Post, error) {
	var post models.Post

	subQuery := p.Db.Model(&models.Post{}).Where("user2_id = ?", userId)
	result := subQuery.Where("id = ?", postId).Find(&post)

	return post, result.Error
}

func (p *PostRepoIpml) CheckUser(userId int) error {
	var user models.User2

	result := p.Db.Where("id = ?", userId).First(&user)
	return result.Error
}

func (p *PostRepoIpml) CreatePost(post models.Post) error {
	result := p.Db.Create(&post)
	return result.Error
}

func (p *PostRepoIpml) CheckUserPost(userId, postId int) error {
	var post models.Post

	subQuery := p.Db.Model(&models.Post{}).Where("user2_id = ?", userId)
	result := subQuery.Where("id = ?", postId).Find(&post)

	return result.Error
}

func (p *PostRepoIpml) UpdatePost(post models.Post) (models.Post, error) {
	result := p.Db.Save(&post)
	return post, result.Error
}

func (p *PostRepoIpml) DeletePost(post models.Post) error {
	result := p.Db.Unscoped().Delete(&post)
	return result.Error
}
