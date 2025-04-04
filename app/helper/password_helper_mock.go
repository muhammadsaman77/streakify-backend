package helper

import "github.com/stretchr/testify/mock"

type PasswordHelperMock struct {
	mock.Mock
}

func (p *PasswordHelperMock) HashPassword(password string) (string, error) {
	args := p.Called(password)
	return args.String(0), args.Error(1)
}

func (p *PasswordHelperMock) CheckPasswordHash(password, hash string) error {
	args := p.Called(password, hash)
	return args.Error(0)
}	