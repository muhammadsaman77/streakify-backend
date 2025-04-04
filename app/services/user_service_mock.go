package services

import (
	"context"

	"github.com/muhammadsaman77/streakify-backend/app/domain/dto"
	"github.com/stretchr/testify/mock"
)



type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) LoginUser(ctx context.Context,loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	args := m.Called(ctx, loginRequest)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.LoginResponse), args.Error(1)
	}
	return nil, args.Error(1)
}