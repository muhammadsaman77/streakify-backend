package services

import (
	"context"
	"testing"

	"github.com/muhammadsaman77/streakify-backend/app/domain/dao"
	"github.com/muhammadsaman77/streakify-backend/app/domain/dto"
	"github.com/muhammadsaman77/streakify-backend/app/helper"
	"github.com/muhammadsaman77/streakify-backend/app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginUser(t *testing.T) {
	mockUserRepository := new(repositories.MockUserRepository)
	mockPasswordHelper := new(helper.PasswordHelperMock)
	mockJWTHelper := new(helper.JWTHelperMock)
	service := NewUserService(mockUserRepository, nil, mockPasswordHelper, mockJWTHelper)

	mockRequest := &dto.LoginRequest{
		Email: "samanmuhammad077@gmail.com",
		Password: "password",
	}
	mockUser := &dao.User{
		ID: 1,
		Email: "samanmuhammad077@gmail.com",
		Password: "hashedPassword",
	}
	mockToken := "token"

	mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything, mockRequest.Email).Return(mockUser, nil)
	mockPasswordHelper.On("CheckPasswordHash", mockRequest.Password, mockUser.Password).Return(nil)
	mockJWTHelper.On("GenerateToken", mockUser.ID, mockUser.Email).Return(mockToken, nil)
	response, err := service.LoginUser(context.Background(), mockRequest)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Token)
	assert.Equal(t, mockToken, response.Token)

	mockUserRepository.AssertExpectations(t)
	mockPasswordHelper.AssertExpectations(t)
	mockJWTHelper.AssertExpectations(t)
}