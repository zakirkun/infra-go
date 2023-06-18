package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zakirkun/infra-go/pkg/auth"
)

func TestJWTImpl_GenerateToken(t *testing.T) {
	data := map[string]interface{}{
		"data":         "farda ayu",
		"authenticate": true,
	}
	token, err := auth.NewJWTImpl("secretKey", 60).GenerateToken(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTImpl_ValidateToken(t *testing.T) {

	data := map[string]interface{}{
		"data":         "farda ayu",
		"authenticate": true,
	}

	token, _ := auth.NewJWTImpl("secretKey", 60).GenerateToken(data)
	valid, err := auth.NewJWTImpl("secretKey", 60).ValidateToken(token)

	assert.NoError(t, err)
	assert.True(t, valid)
}
