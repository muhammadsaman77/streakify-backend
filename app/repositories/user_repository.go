package repositories

import (
	"context"

	"github.com/muhammadsaman77/streakify-backend/app/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (user *dao.User,err error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}


func (r UserRepositoryImpl) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) ( *dao.User, error) {
	var user dao.User
	err := db.Where("email = ?", email).Find(&user).Error
	return &user, err
}