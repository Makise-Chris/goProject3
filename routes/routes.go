package routes

import (
	"goProject3/controllers"
	"goProject3/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signin", controllers.SignIn)
	router.POST("/signup", controllers.SignUp)

	router.POST("/admin/signout", middleware.IsAuthorized(), controllers.SignOut)
	router.POST("/admin/delete/:userid", middleware.IsAuthorized(), controllers.DeleteUser)
	router.GET("/admin", middleware.IsAuthorized(), controllers.AdminIndex)

	router.GET("/user/:userid/:postid", controllers.GetPostByUserId)
	router.PUT("/user/:userid/:postid", middleware.IsAuthorized(), middleware.CheckId(), controllers.UpdatePost)
	router.DELETE("/user/:userid/:postid", middleware.IsAuthorized(), middleware.CheckId(), controllers.DeletePost)
	router.POST("/user/:userid", middleware.IsAuthorized(), middleware.CheckId(), controllers.CreatePost)
	router.GET("/user/:userid", controllers.GetAllPostsByUserId)
	router.POST("/user/signout", middleware.IsAuthorized(), controllers.SignOut)
	router.GET("/user", middleware.IsAuthorized(), controllers.UserIndex)

	router.GET("/post", controllers.GetAllPosts)

	return router
}
