package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ValidatePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))

	return err == nil
}
