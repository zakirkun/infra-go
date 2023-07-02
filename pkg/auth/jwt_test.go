package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zakirkun/infra-go/pkg/auth"
)

func TestJWTImpl_GenerateToken(t *testing.T) {
	data := map[string]interface{}{
		"user":         "farda ayu",
		"authenticate": true,
	}
	token, err := auth.NewJWTImpl("secretKey", 60).GenerateToken(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWTImpl_ValidateToken(t *testing.T) {

	data := map[string]interface{}{
		"user":         "farda ayu",
		"authenticate": true,
	}

	token, _ := auth.NewJWTImpl("secretKey", 60).GenerateToken(data)
	valid, err := auth.NewJWTImpl("secretKey", 60).ValidateToken(token)

	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestJWTImpl_ParseToken(t *testing.T) {

	data := map[string]interface{}{
		"user":         "farda ayu",
		"authenticate": true,
		"email":        "fardaayunurfatika@gmailcom",
	}

	token, _ := auth.NewJWTImpl("secretKey", 60).GenerateToken(data)
	dataPlant, err := auth.NewJWTImpl("secretKey", 60).ParseToken(token)

	assert.NoError(t, err)
	assert.Equal(t, data["user"], dataPlant["user"])
	assert.Equal(t, "fardaayunurfatika@gmailcom", dataPlant["email"])
	assert.True(t, dataPlant["authenticate"].(bool))
	assert.Equal(t, "farda ayu", dataPlant["user"])
}
