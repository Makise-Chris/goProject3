package routes

import (
	"goProject3/controllers"
	"goProject3/middleware"
	"goProject3/repository"
	"goProject3/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepo(db)
	userRepo.Migrate()
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	//userControllerImpl := reflect.ValueOf(userController).Interface().(*controllers.UserControllerImpl)

	postRepo := repository.NewPostRepo(db)
	postRepo.Migrate()
	postService := service.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	router.POST("/signin", userController.SignIn)
	router.POST("/signup", userController.SignUp)

	router.POST("/admin/signout", middleware.IsAuthorized(), userController.SignOut)
	router.POST("/admin/delete/:userid", middleware.IsAuthorized(), userController.DeleteUser)
	router.GET("/admin", middleware.IsAuthorized(), controllers.AdminIndex)

	router.GET("/user/:userid/:postid", postController.GetPostByUserId)
	router.PUT("/user/:userid/:postid", middleware.IsAuthorized(), middleware.CheckId(), postController.UpdatePost)
	router.DELETE("/user/:userid/:postid", middleware.IsAuthorized(), middleware.CheckId(), postController.DeletePost)
	router.POST("/user/:userid", middleware.IsAuthorized(), middleware.CheckId(), postController.CreatePost(userService))
	router.GET("/user/:userid", postController.GetAllPostsByUserId)
	router.POST("/user/signout", middleware.IsAuthorized(), userController.SignOut)
	router.GET("/user", middleware.IsAuthorized(), controllers.UserIndex)

	router.GET("/post", postController.GetAllPosts)

	return router
}
