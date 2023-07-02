package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT interface {
	GenerateToken(data interface{}) (string, error)
	ValidateToken(token interface{}) (bool, error)
}

type JWTImpl struct {
	SignatureKey string
	Expiration   int
}

func (j *JWTImpl) GenerateToken(data interface{}) (string, error) {
	var mySigningKey = []byte(j.SignatureKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["name"] = data
	claims["exp"] = time.Now().Add(time.Duration(j.Expiration) * 24 * 7).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWTImpl) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memastikan bahwa metode penandatanganan token sesuai dengan metode yang digunakan saat pembuatan token.
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

	// Lakukan validasi khusus sesuai dengan kebutuhan aplikasi Anda.
	// Contoh validasi: periksa apakah token telah kedaluwarsa.
	expirationTime := claims["exp"].(float64)
	if time.Now().Unix() > int64(expirationTime) {
		return false, nil
	}

	return true, nil
}
