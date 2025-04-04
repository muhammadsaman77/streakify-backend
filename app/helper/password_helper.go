package helper
type PasswordHelper interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) error
}

type PasswordHelperImpl struct {}

func NewPasswordHelper() PasswordHelper {
	return &PasswordHelperImpl{}
}

func (p *PasswordHelperImpl) HashPassword(password string) (string, error) {
	return password, nil
}

func (p *PasswordHelperImpl) CheckPasswordHash(password, hash string) error {
	return nil
}