package JWT

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var salt = []byte("Sanzhar0GY00bunnyG4y@@@")

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expires":  time.Now().Add(12 * time.Hour).Unix(),
		"issued":   time.Now().Unix(),
	})
	tokenString, _ := token.SignedString(salt)
	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong token")
		}
		return salt, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token if not validated")
}
