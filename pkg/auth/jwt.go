package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT interface {
	GenerateToken(data string) (string, error)
	ValidateToken(token string) (bool, error)
}

type JWTImpl struct {
	SignatureKey string
	Expiration   int
}

func NewJWTImpl(signatureKey string, expiration int) JWT {
	return &JWTImpl{SignatureKey: signatureKey, Expiration: expiration}
}

func (j *JWTImpl) GenerateToken(data string) (string, error) {
	var mySigningKey = []byte(j.SignatureKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["name"] = data

	/**
	-jwt expires in day-
	 for example, if j.Expiration is 20, then the token will expire in 20 days
	**/
	expirationDuration := time.Duration(j.Expiration) * 24 * time.Hour * 7
	expirationTime := time.Now().Add(expirationDuration).Unix()
	claims["exp"] = expirationTime

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTImpl) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SignatureKey), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil
	}

	expirationTime := claims["exp"].(float64)
	if time.Now().Unix() > int64(expirationTime) {
		return false, nil
	}

	return true, nil
}
