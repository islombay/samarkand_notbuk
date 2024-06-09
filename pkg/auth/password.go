package auth

import "golang.org/x/crypto/bcrypt"

func GetHashPassword(s string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}
