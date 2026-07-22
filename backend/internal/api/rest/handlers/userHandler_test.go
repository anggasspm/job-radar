package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct{}

// Mock service
func (m *MockUserService) SignUp(req dto.UserSignup) (*dto.UserSignupResponse, error) {
	return &dto.UserSignupResponse{
		User: dto.UserResponse{
			ID:    1,
			Email: "arif@gmail.com",
		},
		AccessToken:  "mock-access-token",
		RefreshToken: "mock-refresh-token",
	}, nil
}

func (m *MockUserService) Login(req dto.UserLogin) (*dto.UserSigninResponse, error) {
	panic("not implemented")
}

func (m *MockUserService) RefreshToken(token string) (*dto.RefreshTokenResponse, error) {
	panic("not implemented")
}

// HTTP dummy test
func TestRegister_Success(t *testing.T) {
	// turn off logging, maintain cleanness
	gin.SetMode(gin.TestMode)

	// mock service
	mockService := &MockUserService{}

	// mock handler
	handler := NewUserHandler(mockService)

	// mock router
	router := gin.Default()
	router.POST("/register", handler.Register)

	// dummy request
	body := dto.UserSignup{
		UserLogin: dto.UserLogin{
			Email:    "arif@gmail.com",
			Password: "123456",
		},
		Name: "Arif",
	}

	// bind the body
	jsonBody, err := json.Marshal(body)
	assert.NoError(t, err)

	// new request
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// store status code, headher, body, etc
	w := httptest.NewRecorder()
	
	// shoot the request to the router
	router.ServeHTTP(w, req)

	// check the result
	assert.Equal(t, http.StatusCreated, w.Code)

	assert.Contains(
		t,
		w.Body.String(),
		"Register successful",
	)

	// check cookies
	cookies := w.Result().Cookies()

	var accessFound bool
	var refreshFound bool

	// loop all cookies
	for _, cookie := range cookies {
		switch cookie.Name {
		case "access_token":
			accessFound = true
			assert.Equal(t, "mock-access-token", cookie.Value)

		case "refresh_token":
			refreshFound = true
			assert.Equal(t, "mock-refresh-token", cookie.Value)
		}
	}

	assert.True(t, accessFound)
	assert.True(t, refreshFound)
}

func TestRegister_EmptyBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{}

	handler := NewUserHandler(mockService)

	router := gin.Default()
	router.POST("/register", handler.Register)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(""))) // empty body
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

