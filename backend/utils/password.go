package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPwd(pwd string) (string, error) {
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedpwd), nil
}

func CheckPwd(pwd string, hashPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
}
