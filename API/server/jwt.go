package server

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secrete_key")

type user struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	expirationTime := time.Now().Add(300 * time.Minute)
	claims := user{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func IsAValidateToken(tokenString string) bool {
	claims := &user{}
	token, _ := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	return token.Valid
}
