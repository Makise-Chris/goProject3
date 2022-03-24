package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goProject3/models"
	"goProject3/service"
	"goProject3/utils"
)

//go test -coverprofile cover.out
//go tool cover -html cover.out

var password, _ = utils.GeneratehashPassword("Nam12345")

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cases := []struct {
		user           models.User2
		validate       string
		getUserByEmail error
		createUser     error
		want           string
	}{
		{
			user: models.User2{
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			validate: "Nhập thiếu thông tin. ",
			want:     "Nhập thiếu thông tin. ",
		},
		{
			user: models.User2{
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			validate: "Nhập sai định dạng email. ",
			want:     "Nhập sai định dạng email. ",
		},
		{
			user: models.User2{
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "Email already in use",
		},
		{
			getUserByEmail: fmt.Errorf("Cannot get user"),
			createUser:     fmt.Errorf("Cannot create user"),
			want:           "Cannot create user",
		},
		{
			getUserByEmail: fmt.Errorf("Cannot get user"),
			want:           "Sign up successfully!!",
		},
	}

	mockUserService := new(service.MockUserService)
	userController := NewUserController(mockUserService)

	for _, tc := range cases {
		request, _ := json.Marshal(tc.user)

		mockUserService.On("ValidateUser", tc.user).Return(tc.validate)
		mockUserService.On("GetUserByEmail", tc.user.Email).Return(tc.user, tc.getUserByEmail)
		mockUserService.On("CreateUser", mock.AnythingOfType("models.User2")).Return(tc.createUser)

		want := models.JsonResponse{
			Message: tc.want,
		}

		router := gin.Default()

		router.POST("/signup", userController.SignUp)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(request))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)

		mockUserService.ExpectedCalls = nil
	}
}

func TestSignIn(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		auth           models.Authentication
		user           models.User2
		validate       string
		getUserByEmail error
		want           string
	}{
		{
			auth: models.Authentication{
				Email:    "",
				Password: "Nam12345",
			},
			validate: "Nhập thiếu thông tin. ",
			want:     "Nhập thiếu thông tin. ",
		},
		{
			auth: models.Authentication{
				Email:    "nam12345",
				Password: "Nam12345",
			},
			validate: "Nhập sai định dạng email. ",
			want:     "Nhập sai định dạng email. ",
		},
		{
			auth: models.Authentication{
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
			},
			getUserByEmail: fmt.Errorf("Cannot get user"),
			want:           "Email is incorrect",
		},
		{
			auth: models.Authentication{
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
			},
			user: models.User2{
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "Password is incorrect",
		},
		{
			auth: models.Authentication{
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
			},
			user: models.User2{
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: password,
				Role:     "user",
			},
			want: "Sign in successfully!!",
		},
	}

	mockUserService := new(service.MockUserService)
	userController := NewUserController(mockUserService)

	for _, tc := range cases {
		request, _ := json.Marshal(tc.auth)

		mockUserService.On("ValidateAuth", tc.auth).Return(tc.validate)
		mockUserService.On("GetUserByEmail", tc.auth.Email).Return(tc.user, tc.getUserByEmail)

		want := models.JsonResponse{
			Message: tc.want,
		}

		router := gin.Default()

		router.POST("/signin", userController.SignIn)
		req, _ := http.NewRequest("POST", "/signin", bytes.NewReader(request))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)

		mockUserService.ExpectedCalls = nil
	}
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		admin       models.User2
		user        models.User2
		deleteUser  error
		getUserById error
		want        string
	}{
		{
			admin: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "Not authorized",
		},
		{
			admin: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "admin",
			},
			user: models.User2{
				ID: 20,
			},
			getUserById: fmt.Errorf("Cannot get user by ID"),
			want:        "This User does not exist",
		},
		{
			admin: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "admin",
			},
			user: models.User2{
				ID:   20,
				Role: "admin",
			},
			want: "Cannot delete Admin",
		},
		{
			admin: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "admin",
			},
			user: models.User2{
				ID:   20,
				Role: "user",
			},
			want: "Delete User 20 successfully!!",
		},
	}

	for _, tc := range cases {
		validToken, _ := utils.GenerateJWT(tc.admin.Email, tc.admin.Role, int(tc.admin.ID))

		want := models.JsonResponse{
			Message: tc.want,
		}

		mockUserService := new(service.MockUserService)
		userController := NewUserController(mockUserService)

		mockUserService.On("DeleteUser", tc.user).Return(tc.deleteUser)
		mockUserService.On("GetUserById", int(tc.user.ID)).Return(tc.user, tc.getUserById)

		router := gin.Default()

		router.DELETE("/admin/delete/:userid", userController.DeleteUser)
		req, _ := http.NewRequest("DELETE", "/admin/delete/20", nil)
		w := httptest.NewRecorder()

		req.Header.Set("Role", tc.admin.Role)

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: validToken,
		})

		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)
		mockUserService.ExpectedCalls = nil
	}

}
