package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 加密		同bcrypt中的hashpw(pwd, salt)
func PasswordHash(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// 验证		同bcrypt中的checkpw(pwd1, pwd2)
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
