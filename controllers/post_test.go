package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goProject3/models"
	"goProject3/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go test -coverprofile cover.out
//go tool cover -html cover.out

func TestGetAllPost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		posts       []models.Post
		getAllPosts error
		want        string
	}{
		{
			getAllPosts: fmt.Errorf("Cannot paginate"),
			want:        "Cannot paginate",
		},
		{
			posts: []models.Post{
				{
					ID:      1,
					Caption: "Post 1",
					Image:   "post1.png",
				},
				{
					ID:      2,
					Caption: "Post 2",
					Image:   "post2.png",
				},
			},
		},
	}
	mockPostService := new(service.MockPostService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		mockPostService.On("GetAllPosts",
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(tc.posts, tc.getAllPosts)

		want := models.JsonResponse{
			Message: tc.want,
			Data:    tc.posts,
		}

		router := gin.Default()

		router.GET("/post", postController.GetAllPosts)
		req, _ := http.NewRequest("GET", "/post", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)

		mockPostService.ExpectedCalls = nil
	}
}

func TestGetAllPostsByUserId(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		posts               []models.Post
		getAllPostsByUserId error
		want                string
	}{
		{
			getAllPostsByUserId: fmt.Errorf("Cannot paginate"),
			want:                "Cannot paginate",
		},
		{
			posts: []models.Post{
				{
					ID:      1,
					Caption: "Post 1",
					Image:   "post1.png",
				},
				{
					ID:      2,
					Caption: "Post 2",
					Image:   "post2.png",
				},
			},
		},
	}
	mockPostService := new(service.MockPostService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		mockPostService.On("GetAllPostsByUserId", mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(tc.posts, tc.getAllPostsByUserId)

		want := models.JsonResponse{
			Message: tc.want,
			Data:    tc.posts,
		}

		router := gin.Default()

		router.GET("/user/:userid", postController.GetAllPostsByUserId)
		req, _ := http.NewRequest("GET", "/user/10", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)

		mockPostService.ExpectedCalls = nil
	}
}

func TestGetPostByUserId(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		post            models.Post
		getPostByUserId error
		want            string
	}{
		{
			getPostByUserId: fmt.Errorf("This user does not have this post"),
			want:            "This user does not have this post",
		},
		{
			post: models.Post{
				ID:      20,
				Caption: "Post 20",
				Image:   "post20.png",
			},
		},
	}
	mockPostService := new(service.MockPostService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		mockPostService.On("GetPostByUserId", mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(tc.post, tc.getPostByUserId)

		router := gin.Default()

		router.GET("/user/:userid/:postid", postController.GetPostByUserId)
		req, _ := http.NewRequest("GET", "/user/10/20", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, tc.want, got.Message)

		mockPostService.ExpectedCalls = nil
	}
}

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		user         models.User2
		post         models.Post
		validatePost string
		getUserById  error
		want         string
	}{
		{
			user: models.User2{
				ID:   20,
				Role: "admin",
			},
			want: "Admin cannot create post for user",
		},
		{
			getUserById: fmt.Errorf("This User does not exist"),
			want:        "This User does not exist",
		},
		{
			user: models.User2{
				ID:   10,
				Role: "user",
			},
			post: models.Post{
				Image: "post20.png",
			},
			validatePost: "Thiếu caption",
			want:         "Thiếu caption",
		},
		{
			user: models.User2{
				ID:   10,
				Role: "user",
			},
			post: models.Post{
				Caption: "Post 20",
				Image:   "post20.png",
			},
			want: "Create post successfully!!",
		},
	}
	mockPostService := new(service.MockPostService)
	mockUserService := new(service.MockUserService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		request, _ := json.Marshal(tc.post)

		mockUserService.On("GetUserById", mock.AnythingOfType("int")).Return(tc.user, tc.getUserById)
		mockPostService.On("CreatePost", mock.AnythingOfType("models.Post")).Return(nil)
		mockPostService.On("ValidatePost", tc.post).Return(tc.validatePost)

		router := gin.Default()

		router.POST("/user/:userid", postController.CreatePost(mockUserService))
		req, _ := http.NewRequest("POST", "/user/10", bytes.NewReader(request))

		req.Header.Set("Role", tc.user.Role)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, tc.want, got.Message)

		mockPostService.ExpectedCalls = nil
		mockUserService.ExpectedCalls = nil
	}
}

func TestUpdatePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		post          models.Post
		validatePost  string
		checkUserPost error
		want          string
	}{
		{
			checkUserPost: fmt.Errorf("This user does not have this post"),
			want:          "This user does not have this post",
		},
		{
			post: models.Post{
				ID:      20,
				Image:   "post20.png",
				User2ID: 10,
			},
			validatePost: "Thiếu caption",
			want:         "Thiếu caption",
		},
		{
			post: models.Post{
				ID:      20,
				Caption: "Post 20",
				Image:   "post20.png",
				User2ID: 10,
			},
			want: "Update post 20 successfully!!",
		},
	}
	mockPostService := new(service.MockPostService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		request, _ := json.Marshal(tc.post)

		mockPostService.On("UpdatePost", mock.AnythingOfType("models.Post")).Return(nil)
		mockPostService.On("ValidatePost", tc.post).Return(tc.validatePost)
		mockPostService.On("CheckUserPost", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(tc.checkUserPost)

		router := gin.Default()

		router.PUT("/user/:userid/:postid", postController.UpdatePost)
		req, _ := http.NewRequest("PUT", "/user/10/20", bytes.NewReader(request))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, tc.want, got.Message)

		mockPostService.ExpectedCalls = nil
	}
}

func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		checkUserPost error
		want          string
	}{
		{
			checkUserPost: fmt.Errorf("This user does not have this post"),
			want:          "This user does not have this post",
		},
		{
			want: "Delete post 20 successfully!!",
		},
	}
	mockPostService := new(service.MockPostService)
	postController := NewPostController(mockPostService)

	for _, tc := range cases {
		mockPostService.On("DeletePost", mock.AnythingOfType("models.Post")).Return(nil)
		mockPostService.On("CheckUserPost", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(tc.checkUserPost)

		router := gin.Default()

		router.DELETE("/user/:userid/:postid", postController.DeletePost)
		req, _ := http.NewRequest("DELETE", "/user/10/20", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, tc.want, got.Message)

		mockPostService.ExpectedCalls = nil
	}
}
