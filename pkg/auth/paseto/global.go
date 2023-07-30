package paseto

import "github.com/o1egl/paseto/v2"

type IPaseto interface {
	Encrypt(token paseto.JSONToken, footer string) (*string, error)
	Decrypt(_token string) (*paseto.JSONToken, *string, error)
}
