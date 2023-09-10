package services

import (
	config "final-project/config"
	"final-project/model"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(userInfo *model.User) (string, error) {
	config, _ := config.LoadConfig(".")
	jwtSecret := []byte(config.JWT_SECRET_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userInfo.Id,
		"userName": userInfo.Username,
		"email":    userInfo.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func ValidateToken(tokenString string) (map[string]interface{}, error) {
	config, _ := config.LoadConfig(".")
	jwtSecret := []byte(config.JWT_SECRET_KEY)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil

	} else {
		return nil, err
	}
}
