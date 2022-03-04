package controllers

import (
	"goProject3/models"

	"github.com/gin-gonic/gin"
)

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
//@Router /admin/delete/{id} [post]
func DeleteUser(c *gin.Context) {
	if c.Request.Header.Get("Role") != "admin" {
		c.JSON(401, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId := c.Param("userid")

	var dbuser models.User2

	result := models.Connection.Where("id = ?", userId).First(&dbuser)
	if result.Error != nil {
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
		models.Connection.Unscoped().Delete(&dbuser)
		c.JSON(200, gin.H{
			"message": "Delete User " + userId + " successfully!!",
		})
	}
}
