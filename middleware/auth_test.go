package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"goProject3/models"
	"goProject3/utils"
)

//go test -coverprofile cover.out
//go tool cover -html cover.out
func TestIsAuthorized(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		user   models.User2
		signIn bool
		want   string
	}{
		{
			signIn: false,
			want:   "No Token found",
		},
		{
			signIn: true,
			want:   "Not Authorized",
		},
		{
			signIn: true,
			user: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "",
		},
		{
			signIn: true,
			user: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "admin",
			},
			want: "",
		},
	}

	for _, tc := range cases {
		want := models.JsonResponse{
			Message: tc.want,
		}

		router := gin.Default()

		router.GET("/auth", IsAuthorized())

		req, _ := http.NewRequest("GET", "/auth", nil)

		if tc.signIn {
			validToken, _ := utils.GenerateJWT(tc.user.Email, tc.user.Role, int(tc.user.ID))
			cookie := &http.Cookie{
				Name:  "token",
				Value: validToken,
			}
			req.AddCookie(cookie)
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		//reqCookie, _ := req.Cookie("token")
		//assert.Equal(t, cookie, reqCookie)
		assert.Equal(t, want, got)
		assert.Equal(t, req.Header.Get("Role"), tc.user.Role)
	}
}

func TestCheckId(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	cases := []struct {
		user models.User2
		want string
	}{
		{
			user: models.User2{
				ID:       10,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "You cannot do this activity",
		},

		{
			user: models.User2{
				ID:       20,
				Name:     "Nam",
				Email:    "nam12345@gmail.com",
				Password: "Nam12345",
				Role:     "user",
			},
			want: "",
		},
	}

	for _, tc := range cases {
		want := models.JsonResponse{
			Message: tc.want,
		}

		router := gin.Default()
		router.GET("/checkId/:userid", CheckId())

		req, _ := http.NewRequest("GET", "/checkId/20", nil)
		w := httptest.NewRecorder()

		req.Header.Set("Role", tc.user.Role)
		req.Header.Set("ID", strconv.Itoa(int(tc.user.ID)))

		router.ServeHTTP(w, req)

		var got models.JsonResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, want, got)
	}
}
