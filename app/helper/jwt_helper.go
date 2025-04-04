package helper

type JWTHelper interface {
	GenerateToken(id int, email string) (string, error)
	
}

type JWTHelperImpl struct {}

func NewJWTHelper() JWTHelper {
	return &JWTHelperImpl{}
}

func (j *JWTHelperImpl) GenerateToken(id int, email string) (string, error) {
	return "", nil
}
