package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaman77/streakify-backend/app/domain/dto"
	"github.com/muhammadsaman77/streakify-backend/app/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

		mockService := new(services.MockUserService)
		controller := NewUserController(mockService)

		reqBody := dto.LoginRequest{
			Email: "samanmuhammad077@gmail.com",
			Password: "password",
		}
		expectedResponse := dto.LoginResponse{
			Token: "token",
		}
		mockService.On("LoginUser", mock.Anything, &reqBody).Return(&expectedResponse, nil)
		router := gin.Default()
		router.POST("/login", controller.LoginUser)
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w:= httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Success Login")
		assert.Contains(t, w.Body.String(), expectedResponse.Token)
		mockService.AssertExpectations(t)
}

func TestLoginUser_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(services.MockUserService)
	controller := NewUserController(mockService)

	invalidBody := dto.LoginRequest{
		Email: "samanmuhammad077gmail.com",
	}

	router := gin.Default()
	router.POST("/login", controller.LoginUser)
	body, _ := json.Marshal(invalidBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad Request")
	mockService.AssertNotCalled(t, "LoginUser")
}

func TestLoginUser_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(services.MockUserService)
	controller := NewUserController(mockService)

	reqBody := dto.LoginRequest{
		Email: "samanmuhammad077@gmail.com",
		Password: "password",
	}
	mockService.On("LoginUser", mock.Anything, &reqBody).Return(nil,  errors.New("database error"))
	router := gin.Default()
	router.POST("/login", controller.LoginUser)
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Internal Server Error")
	assert.Contains(t, w.Body.String(), "database error")
	mockService.AssertExpectations(t)
}

func TestLoginUser_InvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(services.MockUserService)
	controller := NewUserController(mockService)

	reqBody := dto.LoginRequest{
		Email: "samanmuhammad077@gmail.com",
		Password: "password",
	}
	mockService.On("LoginUser", mock.Anything, &reqBody).Return(nil,  errors.New("invalid credentials"))
	router := gin.Default()
	router.POST("/login", controller.LoginUser)
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
	assert.Contains(t, w.Body.String(), "invalid credentials")
	mockService.AssertExpectations(t)
}