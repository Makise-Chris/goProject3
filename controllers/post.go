package controllers

import (
	"fmt"
	"goProject3/models"
	"goProject3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetAllPosts godoc
//@Summary Lay tat ca cac bai post
//@Description Lay tat ca cac bai post
//@Tags GetAllPosts
//@Accept json
//@Produce json
//@Param page query int false "Page"
//@Param limit query int false "Limit"
//@Param sort query string false "Sort"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /post [get]
func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	queryBuider := models.Connection.Limit(pagination.Limit).Offset(offset).Order("id " + pagination.Sort)
	result := queryBuider.Model(&models.Post{}).Find(&posts)
	//result := models.Connection.Scopes(utils.Paginate(c)).Find(&posts)

	fmt.Println("-------------")
	fmt.Println(pagination)
	fmt.Println(posts)
	fmt.Println(len(posts))
	fmt.Println("-------------")

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Cannot paginate",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}

//GetAllPostsByUserId godoc
//@Summary Lay tat ca cac bai post cua user
//@Description Lay tat ca cac bai post cua user
//@Tags GetAllPostsByUserId
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param page query int false "Page"
//@Param limit query int false "Limit"
//@Param sort query string false "Sort"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid} [get]
func GetAllPostsByUserId(c *gin.Context) {
	var posts []models.Post

	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	userId := c.Param("userid")

	queryBuider := models.Connection.Limit(pagination.Limit).Offset(offset).Order("id " + pagination.Sort)
	result := queryBuider.Model(&models.Post{}).Where("user2_id = ?", userId).Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Cannot paginate",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}

//GetPostByUserId godoc
//@Summary Lay mot bai post cua user
//@Description Lay mot bai post cua user
//@Tags GetPostByUserId
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param  postid path int true "Post ID"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid}/{postid} [get]
func GetPostByUserId(c *gin.Context) {
	userId := c.Param("userid")

	postId := c.Param("postid")

	var post models.Post

	subQuery := models.Connection.Model(&models.Post{}).Where("user2_id = ?", userId)
	subQuery.Where("id = ?", postId).Find(&post)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "This user does not have this post",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

//CreatePost godoc
//@Summary Tao post
//@Description Tao post
//@Tags CreatePost
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param  post body models.PostSwagger true "Create Post"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid} [post]
func CreatePost(c *gin.Context) {
	if c.Request.Header.Get("Role") == "admin" {
		c.JSON(400, gin.H{
			"message": "Admin cannot create post for user",
		})
		return
	}
	userId := c.Param("userid")

	var user models.User2

	result := models.Connection.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "This User does not exist",
		})
		return
	}

	var post models.Post

	c.ShouldBindJSON(&post)
	err := models.Validate.Struct(post)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Thiếu caption",
		})
		return
	}

	post.User2ID, _ = strconv.Atoi(userId)

	models.Connection.Create(&post)

	c.JSON(200, gin.H{
		"message": "Create post successfully!!",
	})
}

//UpdatePost godoc
//@Summary Sua post
//@Description Sua post
//@Tags UpdatePost
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param  post body models.PostSwagger true "Update Post"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid} [put]
func UpdatePost(c *gin.Context) {
	userId := c.Param("userid")

	postId := c.Param("postid")

	var post models.Post

	subQuery := models.Connection.Model(&models.Post{}).Where("user2_id = ?", userId)
	subQuery.Where("id = ?", postId).Find(&post)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "This user does not have this post",
		})
		return
	}

	c.ShouldBindJSON(&post)
	err := models.Validate.Struct(post)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Thiếu caption",
		})
		return
	}

	models.Connection.Save(&post)

	c.JSON(400, gin.H{
		"message": "Update post " + postId + " successfully!",
	})
}

//DeletePost godoc
//@Summary Xoa post
//@Description Xoa post
//@Tags DeletePost
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param  postid path int true "Post ID"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid}/{postid} [delete]
func DeletePost(c *gin.Context) {
	userId := c.Param("userid")

	postId := c.Param("postid")

	var post models.Post

	subQuery := models.Connection.Model(&models.Post{}).Where("user2_id = ?", userId)
	subQuery.Where("id = ?", postId).Find(&post)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "This user does not have this post",
		})
		return
	}

	models.Connection.Unscoped().Delete(&post)
	c.JSON(200, gin.H{
		"message": "Delete post " + postId + " successfully!!",
	})
}
