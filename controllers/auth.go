package controllers

import (
	"goProject3/models"
	"goProject3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
func SignUp(c *gin.Context) {
	var user models.User2
	c.ShouldBindJSON(&user)
	err := models.Validate.Struct(user)
	if err != nil {
		var message string

		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "oneof" {
				message = message + "Nhập sai role (admin hoặc user). "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}

		c.JSON(400, gin.H{
			"message": message,
		})
		return
	}

	var dbuser models.User2
	models.Connection.Where("email = ?", user.Email).First(&dbuser)

	if dbuser.Email != "" {
		c.JSON(400, gin.H{
			"message": "Email already in use",
		})
		return
	}

	user.Password, err = utils.GeneratehashPassword(user.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error in hashing password",
		})
		return
	}

	models.Connection.Create(&user)
	c.JSON(200, gin.H{
		"message": "Sign Up successfully!!",
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
func SignIn(c *gin.Context) {
	var authDetails models.Authentication
	c.ShouldBindJSON(&authDetails)
	err := models.Validate.Struct(authDetails)
	if err != nil {
		var message string

		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "required" {
				message = message + "Nhập thiếu thông tin. "
			}
			if err.ActualTag() == "email" {
				message = message + "Nhập sai định dạng email. "
			}
		}

		c.JSON(400, gin.H{
			"message": message,
		})
		return
	}

	var authUser models.User2
	models.Connection.Where("email = ?", authDetails.Email).First(&authUser)
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
func SignOut(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "token",
		Value: "",
		Path:  "/",
	})
	c.JSON(200, gin.H{
		"message": "Signed Out",
	})
}
