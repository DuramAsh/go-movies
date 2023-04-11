package util

import (
	"github.com/gin-gonic/gin"
)

func HandleRequestError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
