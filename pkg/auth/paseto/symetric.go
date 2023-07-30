package paseto

import (
	"errors"

	"github.com/o1egl/paseto/v2"
)

type pasetoSymmetric struct {
	Key string
}

func NewSymmetric(key string) (IPaseto, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes")
	}

	return &pasetoSymmetric{
		Key: key,
	}, nil
}

func (p pasetoSymmetric) Encrypt(token paseto.JSONToken, footer string) (*string, error) {

	_token, err := paseto.Encrypt([]byte(p.Key), token, footer)
	if err != nil {
		return nil, err
	}

	return &_token, nil
}

func (p pasetoSymmetric) Decrypt(_token string) (*paseto.JSONToken, *string, error) {
	var token paseto.JSONToken
	var footer string

	if err := paseto.Decrypt(_token, []byte(p.Key), &token, &footer); err != nil {
		return nil, nil, err
	}

	return &token, &footer, nil
}
