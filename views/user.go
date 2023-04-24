package views

import (
	"bytes"
	"fmt"
	"github.com/duramash/go-movies/JWT"
	"github.com/duramash/go-movies/models"
	"github.com/duramash/go-movies/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SignUp(ctx *gin.Context) {
	user := models.UserInput{}
	_ = ctx.BindJSON(&user)
	dbUser := models.User{
		Username: user.Username,
		Password: util.HashPassword(&user),
	}
	if err := DB.Create(dbUser).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, dbUser)
}

func Login(ctx *gin.Context) {
	user := models.UserInput{}
	_ = ctx.BindJSON(&user)
	dbUser := models.User{}
	if err := DB.First(&dbUser, "username = ?", user.Username).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusUnauthorized, err)
		return
	}
	if hashedPassword := util.HashPassword(&user); !bytes.Equal(hashedPassword, dbUser.Password) {
		util.HandleRequestError(ctx, http.StatusUnauthorized, fmt.Errorf("wrong password"))
		return
	}
	token, _ := JWT.GenerateToken(user.Username)
	tokens := make(map[string]interface{})
	tokens["access"] = token
	tokens["refresh"] = "to be implemented"
	ctx.JSON(http.StatusOK, tokens)
}

func IdentifyUser(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		util.HandleRequestError(ctx, http.StatusUnauthorized, fmt.Errorf("no auth header"))
		ctx.Abort()
		return
	}
	headerInSlice := strings.Split(header, " ")
	if len(headerInSlice) != 2 {
		util.HandleRequestError(ctx, http.StatusBadRequest, fmt.Errorf("wrong header"))
		ctx.Abort()
		return
	}

	_, err := JWT.ParseToken(headerInSlice[1])
	if err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		ctx.Abort()
		return
	}
}
