package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), 15)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func VerifyPassword(inputPass, storedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPass), []byte(inputPass))
}
