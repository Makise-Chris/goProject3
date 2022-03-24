package controllers

import (
	"goProject3/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	DeleteUser(c *gin.Context)
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(s service.UserService) UserController {
	return &UserControllerImpl{
		UserService: s,
	}
}

//DeleteUser godoc
//@Summary  Xoa User
//@Description Xoa User
//@Tags DeleteUser
//@Accept json
//@Produce json
//@Param  userid path int true "User ID"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Failure 401 {object} models.JsonResponse
//@Router /admin/delete/{userid} [post]
func (u *UserControllerImpl) DeleteUser(c *gin.Context) {
	if c.Request.Header.Get("Role") != "admin" {
		c.JSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId := c.Param("userid")
	userIdInt, _ := strconv.Atoi(userId)

	dbuser, err := u.UserService.GetUserById(userIdInt)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "This User does not exist",
		})
		return
	}

	if dbuser.Role == "admin" {
		c.JSON(400, gin.H{
			"message": "Cannot delete Admin",
		})
	} else {
		u.UserService.DeleteUser(dbuser)
		c.JSON(200, gin.H{
			"message": "Delete User " + userId + " successfully!!",
		})
	}
}
