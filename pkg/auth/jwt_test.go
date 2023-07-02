package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zakirkun/infra-go/pkg/auth"
)

func TestJWTImpl_GenerateToken(t *testing.T) {
	jwtImpl := &auth.JWTImpl{
		SignatureKey: "secretKey",
		Expiration:   60,
	}

	data := "Farda Ayu"
	token, err := jwtImpl.GenerateToken(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTImpl_ValidateToken(t *testing.T) {
	jwtImpl := &auth.JWTImpl{
		SignatureKey: "secretKey",
		Expiration:   60,
	}

	data := "Farda Ayu"
	token, _ := jwtImpl.GenerateToken(data)

	valid, err := jwtImpl.ValidateToken(token)

	assert.NoError(t, err)
	assert.True(t, valid)
}
