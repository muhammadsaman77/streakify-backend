package services

import (
	"context"

	"github.com/muhammadsaman77/streakify-backend/app/domain/dto"
	"github.com/muhammadsaman77/streakify-backend/app/helper"
	"github.com/muhammadsaman77/streakify-backend/app/repositories"
	"gorm.io/gorm"
)


type UserService interface {
	LoginUser(ctx context.Context,loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB *gorm.DB
	PasswordHelper helper.PasswordHelper
	JWTHelper helper.JWTHelper
	
}

func NewUserService(userRepository repositories.UserRepository, db *gorm.DB, passwordHelper helper.PasswordHelper, jwtHelper helper.JWTHelper) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: db,
		PasswordHelper: passwordHelper,
		JWTHelper: jwtHelper,
	}
}

func (service UserServiceImpl) LoginUser(ctx context.Context,loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err :=  service.UserRepository.GetUserByEmail(ctx, service.DB, loginRequest.Email)
	if err != nil {
		return nil, err
	}

	err = service.PasswordHelper.CheckPasswordHash(loginRequest.Password, user.Password)
	if err != nil {
		return nil, err
	}
	token, err := service.JWTHelper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token},nil
}
