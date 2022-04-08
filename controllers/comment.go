package controllers

import (
	"goProject3/elasticsearch"
	"goProject3/models"
	"goProject3/service"
	"goProject3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentControlller interface {
	GetAllCommentsByPostId(c *gin.Context)
	GetAllCommentsByPostIdAndUserId(c *gin.Context)
	CreateComment(u service.UserService, p service.PostService) gin.HandlerFunc
	UpdateComment(u service.UserService) gin.HandlerFunc
	DeleteComment(c *gin.Context)
}

type CommentControlllerImpl struct {
	CommentService service.CommentService
	CommentES      elasticsearch.CommentES
}

func NewCommentController(s service.CommentService, e elasticsearch.CommentES) CommentControlller {
	return &CommentControlllerImpl{
		CommentService: s,
		CommentES:      e,
	}
}

func (cmt *CommentControlllerImpl) GetAllCommentsByPostId(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	postId := c.Param("postid")
	postIdInt, _ := strconv.Atoi(postId)

	comments, err := cmt.CommentService.GetAllCommentsByPostId(postIdInt, pagination.Limit, offset, pagination.Sort)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot paginate",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": comments,
	})
}

func (cmt *CommentControlllerImpl) GetAllCommentsByPostIdAndUserId(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	postId := c.Param("postid")
	postIdInt, _ := strconv.Atoi(postId)

	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	comments, err := cmt.CommentService.GetAllCommentsByPostIdAndUserId(postIdInt, userIdInt, pagination.Limit, offset, pagination.Sort)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot paginate",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": comments,
	})
}

func (cmt *CommentControlllerImpl) CreateComment(u service.UserService, p service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Role") == "admin" {
			c.JSON(400, gin.H{
				"message": "Admin cannot create post for user",
			})
			return
		}

		postId := c.Param("postid")
		postIdInt, _ := strconv.Atoi(postId)

		_, err := p.GetPostById(postIdInt)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "This Post does not exist",
			})
			return
		}

		userId := c.Param("userid")
		userIdInt, _ := strconv.Atoi(userId)

		user, err := u.GetUserById(userIdInt)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "This User does not exist",
			})
			return
		}

		var comment models.Comment

		c.ShouldBindJSON(&comment)

		comment.PostID, _ = strconv.Atoi(postId)
		comment.User2ID, _ = strconv.Atoi(userId)
		comment.User = user

		dbcomment, _ := cmt.CommentService.UpdateComment(comment)
		cmt.CommentES.CreateComment(c.Request.Context(), dbcomment)

		c.JSON(200, gin.H{
			"message": "Create comment for post " + postId + " successfully!!",
		})
	}
}

func (cmt *CommentControlllerImpl) UpdateComment(u service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		userIdInt, _ := strconv.Atoi(userId)

		commentId := c.Param("commentid")
		commentIdInt, _ := strconv.Atoi(commentId)

		err := cmt.CommentService.CheckUserComment(userIdInt, commentIdInt)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "This user does not have this comment",
			})
			return
		}

		user, err := u.GetUserById(userIdInt)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "This User does not exist",
			})
			return
		}

		var comment models.Comment
		c.ShouldBindJSON(&comment)

		comment.ID = uint(commentIdInt)
		comment.User2ID = userIdInt
		comment.User = user
		cmt.CommentService.UpdateComment(comment)
		cmt.CommentES.CreateComment(c.Request.Context(), comment)

		c.JSON(200, gin.H{
			"message": "Update comment " + commentId + " successfully!!",
		})
	}
}

func (cmt *CommentControlllerImpl) DeleteComment(c *gin.Context) {
	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	commentId := c.Param("commentid")
	commentIdInt, _ := strconv.Atoi(commentId)

	err := cmt.CommentService.CheckUserComment(userIdInt, commentIdInt)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "This user does not have this comment",
		})
		return
	}

	var comment models.Comment
	comment.ID = uint(commentIdInt)
	cmt.CommentService.DeleteComment(comment)
	cmt.CommentES.DeleteComment(c.Request.Context(), commentIdInt)
	c.JSON(200, gin.H{
		"message": "Delete comment " + commentId + " successfully!!",
	})
}
