package transformer

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

func Encode(text string, secret string, seed []byte, padding string) string {
	key := strings.ToLower(secret + padding)
	block, err := aes.NewCipher([]byte(key[:16]))
	if err != nil {
		return ""
	}

	raw := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, seed)
	encoded := make([]byte, len(raw))
	cfb.XORKeyStream(encoded, raw)
	return base64.StdEncoding.EncodeToString(encoded)
}

func Decode(text string, secret string, seed []byte, padding string) string {
	key := strings.ToLower(secret + padding)
	block, err := aes.NewCipher([]byte(key[:16]))
	if err != nil {
		return ""
	}

	encoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err, 2)
		return ""
	}

	cfb := cipher.NewCFBDecrypter(block, seed)
	raw := make([]byte, len(encoded))
	cfb.XORKeyStream(raw, encoded)
	return string(raw)
}
