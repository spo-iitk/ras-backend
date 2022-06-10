package auth

import (
	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
