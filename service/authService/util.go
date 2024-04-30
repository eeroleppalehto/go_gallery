package authservice

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwBytes []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pwBytes, 14)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ComparePassowrds(hashedBytes []byte, pw string) bool {
	err := bcrypt.CompareHashAndPassword(hashedBytes, []byte(pw))
	isValid := err == nil
	return isValid
}
