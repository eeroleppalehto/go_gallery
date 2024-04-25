package authservice

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	Username string `json:"name"`
	jwt.RegisteredClaims
}

type AuthService struct{}

func (a *AuthService) HashPassword(pwBytes []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pwBytes, 14)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (a *AuthService) ComparePassowrds(hashedBytes []byte, pw string) bool {
	err := bcrypt.CompareHashAndPassword(hashedBytes, []byte(pw))
	isValid := err == nil
	return isValid
}

func (a *AuthService) NewJwtCustomClaims(username string) (string, error) {
	claims := &JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
