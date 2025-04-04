package repositories

import (
	"context"

	"github.com/muhammadsaman77/streakify-backend/app/domain/dao"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}


func (m *MockUserRepository) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*dao.User, error){
	args := m.Called(ctx, db, email)
	if args.Get(0) != nil {
		return args.Get(0).(*dao.User), args.Error(1)
	}
	return nil, args.Error(1)
}

