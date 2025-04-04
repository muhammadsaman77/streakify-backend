package repositories

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/muhammadsaman77/streakify-backend/app/domain/dao"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhammadsaman77/streakify-backend/config"
	"gorm.io/gorm"
)

func TestGetUserByEmail(t *testing.T) {
	gormDB, mock, err := config.SetupMockDB()
	require.NoError(t, err)


	expectedUser := &dao.User{
		ID:     1,
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
		Timezone: faker.Timezone(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).WithArgs(expectedUser.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password", "timezone", "created_at", "updated_at", "deleted_at"}).
			AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Email, expectedUser.Password, expectedUser.Timezone, expectedUser.CreatedAt, expectedUser.UpdatedAt, expectedUser.DeletedAt))

	
	repo := NewUserRepository()
	user,err :=repo.GetUserByEmail(context.Background(),gormDB, expectedUser.Email)
	assert.NoError(t, err)
	assert.NotNil(t,user)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Username, user.Username)
	require.NoError(t, mock.ExpectationsWereMet())
}

