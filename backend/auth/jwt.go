package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	ttl = 1000000 * time.Second
	hmacSecret = "not-secure"
)

func GetToken(username string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp": time.Now().UTC().Add(ttl).Unix(),
	}).SignedString([]byte(hmacSecret))
}

func GetUsername(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}
