package routes

import (
	"fmt"
	"goProject3/controllers"
	"goProject3/elasticsearch"
	"goProject3/middleware"
	"goProject3/repository"
	"goProject3/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	es "github.com/olivere/elastic/v7"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	client, err := es.NewClient(
		es.SetURL("https://localhost:9200"),
		es.SetBasicAuth("elastic", "ze2TijQ9ai7k-2c=0tNx"),
		es.SetSniff(false),
	)

	if err != nil {
		fmt.Println("Cannot create ES client")
	}

	UserES := elasticsearch.NewUserES(client)
	PostES := elasticsearch.NewPostES(client)
	CommentES := elasticsearch.NewCommentES(client)

	userRepo := repository.NewUserRepo(db)
	userRepo.Migrate()
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, *UserES)

	postRepo := repository.NewPostRepo(db)
	postRepo.Migrate()
	postService := service.NewPostService(postRepo)
	postController := controllers.NewPostController(postService, *PostES)

	commentRepo := repository.NewCommentRepo(db)
	commentRepo.Migrate()
	commentService := service.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService, *CommentES)

	router.POST("/signin", userController.SignIn)
	router.POST("/signup", userController.SignUp)

	router.POST("/admin/signout", middleware.IsAuthorized(), userController.SignOut)
	router.POST("/admin/delete/:userid", middleware.IsAuthorized(), userController.DeleteUser)
	router.GET("/admin", middleware.IsAuthorized(), controllers.AdminIndex)

	router.GET("/user/:userid/post/:postid", postController.GetPostByUserId)
	router.PUT("/user/:userid/post/:postid", middleware.IsAuthorized(), middleware.CheckId(), postController.UpdatePost)
	router.DELETE("/user/:userid/post/:postid", middleware.IsAuthorized(), middleware.CheckId(), postController.DeletePost)
	router.POST("/user/:userid/post/:postid", middleware.IsAuthorized(), middleware.CheckId(), commentController.CreateComment(userService, postService))

	router.PUT("/user/:userid/post/:postid/comment/:commentid", middleware.IsAuthorized(), middleware.CheckId(), commentController.UpdateComment(userService))
	router.DELETE("/user/:userid/post/:postid/comment/:commentid", middleware.IsAuthorized(), middleware.CheckId(), commentController.DeleteComment)

	router.POST("/user/:userid", middleware.IsAuthorized(), middleware.CheckId(), postController.CreatePost(userService))
	router.GET("/user/:userid", postController.GetAllPostsByUserId)

	router.POST("/user/signout", middleware.IsAuthorized(), userController.SignOut)
	router.GET("/user/search/:keyword", userController.SearchUser)
	router.GET("/user", middleware.IsAuthorized(), controllers.UserIndex)

	router.GET("/post", postController.GetAllPosts)
	router.GET("/post/:postid", commentController.GetAllCommentsByPostId)
	router.GET("/post/search/:keyword", postController.SearchPost(*CommentES))
	router.GET("/post/:postid/user/:userid", commentController.GetAllCommentsByPostIdAndUserId)

	return router
}
