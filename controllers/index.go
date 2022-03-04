package controllers

import (
	"github.com/gin-gonic/gin"
)

//AdminIndex godoc
//@Summary Lay trang chu Admin
//@Description Lay trang chu Admin
//@Tags AdminIndex
//@Accept json
//@Produce json
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Failure 401 {object} models.JsonResponse
//@Router /admin [get]
func AdminIndex(c *gin.Context) {
	if c.Request.Header.Get("Role") != "admin" {
		c.JSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Welcome, Admin",
	})
}

//UserIndex godoc
//@Summary Lay trang chu User
//@Description Lay trang chu User
//@Tags UserIndex
//@Accept json
//@Produce json
//@Success 200 {object} models.JsonResponse
//@Failure 400 {object} models.JsonResponse
//@Failure 401 {object} models.JsonResponse
//@Router /user [get]
func UserIndex(c *gin.Context) {
	if c.Request.Header.Get("Role") != "user" {
		c.JSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Welcome, User",
	})
}
