package repos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func GenerateToken(id string) string {
	secretKey, err := hex.DecodeString(os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	expirationTime := time.Now().Add(1 * time.Hour)

	message := []byte(id + expirationTime.Format(time.RFC3339))

	hmac := hmac.New(sha256.New, secretKey)
	hmac.Write(message)
	signature := hmac.Sum(nil)

	token := base64.URLEncoding.EncodeToString(append(signature, []byte(expirationTime.Format(time.RFC3339))...))

	return token
}

//Em construção
