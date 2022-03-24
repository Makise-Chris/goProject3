package controllers

import (
	"goProject3/models"
	"goProject3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SignUp godoc
//@Summary Dang ky
//@Description Dang ky
//@Tags SignUp
//@Accept json
//@Produce json
//@Param  signup body models.User2SignUp true "Sign Up"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /signup [post]
func (u *UserControllerImpl) SignUp(c *gin.Context) {
	var user models.User2
	c.ShouldBindJSON(&user)
	message := u.UserService.ValidateUser(user)
	if message != "" {
		c.JSON(400, gin.H{
			"message": message,
		})
		return
	}

	dbuser, _ := u.UserService.GetUserByEmail(user.Email)

	if dbuser.Email != "" {
		c.JSON(400, gin.H{
			"message": "Email already in use",
		})
		return
	}

	user.Password, _ = utils.GeneratehashPassword(user.Password)
	err := u.UserService.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Sign up successfully!!",
	})
}

//SignIn godoc
//@Summary Dang nhap
//@Description Dang nhap
//@Tags SignIn
//@Accept json
//@Produce json
//@Param  signin body models.Authentication true "Sign In"
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Router /signin [post]
func (u *UserControllerImpl) SignIn(c *gin.Context) {
	var authDetails models.Authentication
	c.ShouldBindJSON(&authDetails)
	message := u.UserService.ValidateAuth(authDetails)
	if message != "" {
		c.JSON(400, gin.H{
			"message": message,
		})
		return
	}

	authUser, _ := u.UserService.GetUserByEmail(authDetails.Email)
	if authUser.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is incorrect",
			"id":      authUser.ID,
		})
	} else {
		check := utils.CheckPasswordHash(authDetails.Password, authUser.Password)
		if !check {
			c.JSON(400, gin.H{
				"message": "Password is incorrect",
			})
			return
		}

		validToken, err := utils.GenerateJWT(authUser.Email, authUser.Role, int(authUser.ID))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to generate token",
			})
			return
		}

		var token models.Token
		token.Email = authUser.Email
		token.Role = authUser.Role
		token.TokenString = validToken

		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "token",
			Value: validToken,
		})

		c.JSON(200, gin.H{
			"message": "Sign in successfully!!",
			"email":   token.Email,
			"role":    token.Role,
			"token":   token.TokenString,
		})
	}
}

//SignOut godoc
//@Summary Dang xuat
//@Description Dang xuat
//@Tags SignOut
//@Accept json
//@Produce json
//@Success 200 {object} models.JsonResponse
//@Router /user/signout [post]
//@Router /admin/signout [post]
func (u *UserControllerImpl) SignOut(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "token",
		Value: "",
		Path:  "/",
	})
	c.JSON(200, gin.H{
		"message": "Signed Out",
	})
}
