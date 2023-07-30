package paseto

import (
	"fmt"
	"testing"
	"time"

	"github.com/o1egl/paseto/v2"
)

func TestEncrypt(t *testing.T) {
	symmetricPaseto, _ := NewSymmetric("PrU5AbXJawKJIUOJFmd4f6ZwmifLvvoF")
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
	token.Set("email", "zakir@bekasi.dev")
	token.Set("name", "Im Zakir")

	// Encrypt token with footer Kitabisa.com
	encryptedToken, _ := symmetricPaseto.Encrypt(token, "indonesia.com")
	fmt.Println(*encryptedToken)
}

func TestDecrypt(t *testing.T) {

	symmetricPaseto, _ := NewSymmetric("PrU5AbXJawKJIUOJFmd4f6ZwmifLvvoF")

	encToken := "v2.local.wFYxli_i41WlKGZM9UHBcxxLV8R3naLeg3vrC7BniIixA72JftmuLIXEmOecEB2IARjWbe_0XckEIIDXpG-9jDITy25OzsaIoOerHhMYY3QYXO1Xxpf3mP88Dk6HPd1mX73eSMbf0rv4eDPtIkwzBEXngnTHwWVTtcIWoA9vhr07k7DasJ56vvGrfaA7MDEFi21XnZhsnYPe7gnMceYFupSFA9kjVA2HbpRx9Rec20Eu1ZfOOJlyneKMwB26--3_y8vX97PqulvhAbyrZ-1BwzRMt47V55-cd5Mr6tWkjgXnP9jZv0mtVQ0yigCzXd33vtEf4MVjjJABNnvovANVbAZW8dGJGFuwzlddLKB1tJ7kzaxRacUoQTaX7vlBGR-fC_omeir1TN-ASAVxmXGjeMUfyyjMSUQPk0jigdCC1CPdX4wtu-jyYQ.aW5kb25lc2lhLmNvbQ"
	decToken, footer, err := symmetricPaseto.Decrypt(encToken)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Decrypted token: %+v\n", *decToken)
	fmt.Println("Footer:", *footer)
}
