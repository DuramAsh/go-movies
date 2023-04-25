package util

import (
	"crypto/sha256"
	"github.com/duramash/go-movies/models"
	"github.com/gin-gonic/gin"
)

func HandleRequestError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func HashPassword(user *models.UserInput) []byte {
	hashFunction := sha256.New()
	hashFunction.Write([]byte(user.Password))
	return hashFunction.Sum(nil)
}

func VerifyPassword(password string) bool {
	return len(password) >= 8 && HasUpperCase(password) && HasDigits(password) && HasSpecialChars(password)
}

func HasUpperCase(s string) bool {
	for _, el := range s {
		if 'A' <= el && el <= 'Z' {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, el := range s {
		if '0' <= el && el <= '9' {
			return true
		}
	}
	return false
}

func HasSpecialChars(s string) bool {
	chars := []rune{'.', '_', ',', '?', '!', '%', '$', '-'}
	for _, el := range s {
		for _, c := range chars {
			if el == c {
				return true
			}
		}
	}
	return false
}
