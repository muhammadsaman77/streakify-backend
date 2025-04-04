package helper

import "github.com/stretchr/testify/mock"

type JWTHelperMock struct {
	mock.Mock
}

func (j *JWTHelperMock) GenerateToken(id int, email string) (string, error) {
	args := j.Called(id,email)
	return args.String(0), args.Error(1)
}