package paseto

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/o1egl/paseto/v2"
)

func TestAsymentricEncrypt(t *testing.T) {
	pub := "ee75e7423709a45dfdcd611a9170e6e5a47d6dd0e49d080caf5187dfb3fd6b2a"
	priv := "3e468998b4bde952b86a89a546c008550c7c774d9e360da2ec7df524892f3468ee75e7423709a45dfdcd611a9170e6e5a47d6dd0e49d080caf5187dfb3fd6b2a"

	ast, _ := NewAsymetric(pub, priv)
	now := time.Now()

	// Standard token data
	token := paseto.JSONToken{
		Audience:   "Kitabisa services",
		Issuer:     "Kulonuwun",
		Jti:        "706cbfce-c031-4a44-815e-030f963f7d4e",
		Subject:    "cac2ee7e-70d0-4220-badd-7b5695f53ad8",
		Expiration: now.Add(24 * time.Hour),
		IssuedAt:   now,
		NotBefore:  now,
	}

	// Set your custom data here
	token.Set("email", "budi@kitabisa.com")
	token.Set("name", "Budi Ariyanto")

	// Encrypt token with footer Kitabisa.com
	encryptedToken, _ := ast.Encrypt(token, "Kitabisa.com")
	fmt.Println(*encryptedToken)
}

func TestGeneratePairFile(t *testing.T) {

	pubkey, privKey, _ := ed25519.GenerateKey(nil)
	pub := hex.EncodeToString(pubkey)
	priv := hex.EncodeToString(privKey)

	fmt.Println("Public key:", pub)
	fmt.Println("Private key:", priv)
}
