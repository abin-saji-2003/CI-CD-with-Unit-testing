package delivery_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"testing/internal/api"
	"testing/internal/core"
	"testing/internal/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() (*gin.Engine, *mocks.MockUserUseCase) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mocks.MockUserUseCase)
	r := api.SetupRouter(mockUseCase)

	return r, mockUseCase
}

func TestAddUser(t *testing.T) {
	router, mockUseCase := setupRouter()

	mockUser := &core.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	mockUseCase.On("AddUser", mock.Anything).Return(nil)

	requestBody, _ := json.Marshal(mockUser)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	router, mockUseCase := setupRouter()

	mockUser := &core.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	mockUseCase.On("GetUser", uint(1)).Return(mockUser, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var returnedUser core.User
	json.Unmarshal(resp.Body.Bytes(), &returnedUser)
	assert.Equal(t, mockUser.Name, returnedUser.Name)
}
