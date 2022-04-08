package controllers

import (
	"goProject3/elasticsearch"
	"goProject3/models"
	"goProject3/service"
	"goProject3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostControlller interface {
	GetAllPosts(c *gin.Context)
	GetAllPostsByUserId(c *gin.Context)
	GetPostByUserId(c *gin.Context)
	CreatePost(u service.UserService) gin.HandlerFunc
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	SearchPost(cmt elasticsearch.CommentES) gin.HandlerFunc
}

type PostControlllerImpl struct {
	PostService service.PostService
	PostES      elasticsearch.PostES
}

func NewPostController(s service.PostService, e elasticsearch.PostES) PostControlller {
	return &PostControlllerImpl{
		PostService: s,
		PostES:      e,
	}
}

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
func (p *PostControlllerImpl) GetAllPosts(c *gin.Context) {
	var posts []models.Post

	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	posts, err := p.PostService.GetAllPosts(pagination.Limit, offset, pagination.Sort)
	//result := models.Connection.Scopes(utils.Paginate(c)).Find(&posts)

	if err != nil {
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
func (p *PostControlllerImpl) GetAllPostsByUserId(c *gin.Context) {
	var posts []models.Post

	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	posts, err := p.PostService.GetAllPostsByUserId(userIdInt, pagination.Limit, offset, pagination.Sort)

	if err != nil {
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
//@Router /user/{userid}/post/{postid} [get]
func (p *PostControlllerImpl) GetPostByUserId(c *gin.Context) {
	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	postId := c.Param("postid")
	posiIdInt, _ := strconv.Atoi(postId)

	post, _ := p.PostService.GetPostByUserId(userIdInt, posiIdInt)
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
func (p *PostControlllerImpl) CreatePost(u service.UserService) gin.HandlerFunc {
	create := func(c *gin.Context) {
		if c.Request.Header.Get("Role") == "admin" {
			c.JSON(400, gin.H{
				"message": "Admin cannot create post for user",
			})
			return
		}
		userId := c.Param("userid")
		userIdInt, _ := strconv.Atoi(userId)

		_, err := u.GetUserById(userIdInt)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "This User does not exist",
			})
			return
		}

		var post models.Post

		c.ShouldBindJSON(&post)

		message := p.PostService.ValidatePost(post)
		if message != "" {
			c.JSON(400, gin.H{
				"message": message,
			})
			return
		}

		post.User2ID, _ = strconv.Atoi(userId)

		dbposts, _ := p.PostService.UpdatePost(post)
		p.PostES.CreatePost(c.Request.Context(), dbposts)

		c.JSON(200, gin.H{
			"message": "Create post successfully!!",
		})
	}
	return gin.HandlerFunc(create)
}

//UpdatePost godoc
//@Summary Sua post
//@Description Sua post
//@Tags UpdatePost
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Param  postid path int true "Post ID"
//@Param  post body models.PostSwagger true "Update Post"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /user/{userid}/post/{postid} [put]
func (p *PostControlllerImpl) UpdatePost(c *gin.Context) {
	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	postId := c.Param("postid")
	postIdInt, _ := strconv.Atoi(postId)

	err := p.PostService.CheckUserPost(userIdInt, postIdInt)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "This user does not have this post",
		})
		return
	}

	var post models.Post
	c.ShouldBindJSON(&post)
	message := p.PostService.ValidatePost(post)
	if message != "" {
		c.JSON(400, gin.H{
			"message": message,
		})
		return
	}

	post.ID = uint(postIdInt)
	post.User2ID = userIdInt
	p.PostService.UpdatePost(post)
	p.PostES.CreatePost(c.Request.Context(), post)

	c.JSON(200, gin.H{
		"message": "Update post " + postId + " successfully!!",
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
//@Router /user/{userid}/post/{postid} [delete]
func (p *PostControlllerImpl) DeletePost(c *gin.Context) {
	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	postId := c.Param("postid")
	postIdInt, _ := strconv.Atoi(postId)

	err := p.PostService.CheckUserPost(userIdInt, postIdInt)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "This user does not have this post",
		})
		return
	}

	var post models.Post
	post.ID = uint(postIdInt)
	p.PostService.DeletePost(post)
	p.PostES.DeletePost(c.Request.Context(), postIdInt)
	c.JSON(200, gin.H{
		"message": "Delete post " + postId + " successfully!!",
	})
}

func (p *PostControlllerImpl) SearchPost(cmt elasticsearch.CommentES) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("keyword")

		posts, err := p.PostES.SearchPost(c.Request.Context(), query)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		comments, err := cmt.SearchComment(c.Request.Context(), query)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"posts":    posts,
			"comments": comments,
		})
	}
}
