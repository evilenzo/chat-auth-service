package main

import (
	"github.com/matthewhartstonge/argon2"
	"golang.org/x/crypto/bcrypt"
)

func check(where string, err error) {

}

func GetToken(codeword []byte) string {
	hash, err := bcrypt.GenerateFromPassword(codeword, bcrypt.DefaultCost)
	check("token generation", err)

	return string(hash)
}

func VerifyToken(codeword []byte, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(token), codeword)
	if err != nil {
		return false
	}

	return true
}

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	check("argon password hashing", err)

	return string(encoded)
}

func VerifyPassword(password string, hash string) bool {
	matches, err := argon2.VerifyEncoded([]byte(password), []byte(hash))
	check("password hash verification", err)
	return matches
}
