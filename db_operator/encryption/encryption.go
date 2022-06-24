package encryption

import (
	"github.com/matthewhartstonge/argon2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GetToken(codeword []byte) string {
	hash, err := bcrypt.GenerateFromPassword(codeword, bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error while token generation: ", err)
	}

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
	if err != nil {
		log.Error("Error while hash encoding: ", err)
	}

	return string(encoded)
}

func VerifyPassword(password string, hash string) bool {
	matches, err := argon2.VerifyEncoded([]byte(password), []byte(hash))
	if err != nil {
		log.Error("Error while hash verification: ", err)
	}

	return matches
}
