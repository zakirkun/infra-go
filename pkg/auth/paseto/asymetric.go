package paseto

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/o1egl/paseto/v2"
)

type pasetoAsymetric struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func NewAsymetric(pub, priv string) (IPaseto, error) {
	_priv, err := hex.DecodeString(priv)
	if err != nil {
		return nil, err
	}

	privateKey := ed25519.PrivateKey(_priv)

	_pub, err := hex.DecodeString(pub)
	if err != nil {
		return nil, err
	}

	publicKey := ed25519.PublicKey(_pub)

	return pasetoAsymetric{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func (p pasetoAsymetric) Encrypt(token paseto.JSONToken, footer string) (*string, error) {

	_token, err := paseto.Sign(p.PrivateKey, token, footer)
	if err != nil {
		return nil, err
	}

	return &_token, nil
}

func (p pasetoAsymetric) Decrypt(_token string) (*paseto.JSONToken, *string, error) {

	var token paseto.JSONToken
	var footer string

	if err := paseto.Verify(_token, p.PublicKey, &token, &footer); err != nil {
		return nil, nil, err
	}

	return &token, &footer, nil
}
