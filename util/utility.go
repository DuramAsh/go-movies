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
