package middleware

import (
	"fmt"
	"net/http"

	"goProject3/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(401, gin.H{
					"message": "No Token found",
				})
				c.Abort()
				return
			}
			c.JSON(400, gin.H{
				"message": "Something went wrong when get cookie",
			})
			c.Abort()
			return
		}
		tokenStr := cookie.Value

		var mySigningKey = []byte(models.Secretkey)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Your Token has been expired",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["id"]
			c.Request.Header.Set("ID", fmt.Sprint(id))

			if claims["role"] == "admin" {
				c.Request.Header.Set("Role", "admin")
				return

			} else if claims["role"] == "user" {
				c.Request.Header.Set("Role", "user")
				return
			}
		}

		c.JSON(400, gin.H{
			"message": "Not Authorized",
		})
	}
}

func CheckId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Role") == "user" && c.Request.Header.Get("ID") != c.Param("userid") {
			c.JSON(400, gin.H{
				"message": "You cannot do this activity",
			})
			c.Abort()
			return
		}
	}
}
