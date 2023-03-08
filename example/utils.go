package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/htmlindex"
)

func GetLocalApiAddr(api string) string {
	return fmt.Sprintf("%s%s%s", BASE_API_HOST, DEFAULT_PORT, api)
}

func GetDockerApiAddr(api string) string {
	return fmt.Sprintf("%s%s%s", BASE_API_IN_DOCKER, DEFAULT_PORT, api)
}

func GetVar(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func GetEnv(key string) string {
	return GetVar(os.Getenv(key))
}

func Encode(text, secret, seed, padding string) string {
	key := strings.ToLower(secret + padding)
	block, err := aes.NewCipher([]byte(key[:16]))
	if err != nil {
		return ""
	}
	raw := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(seed))
	encoded := make([]byte, len(raw))
	cfb.XORKeyStream(encoded, raw)
	return base64.StdEncoding.EncodeToString(encoded)
}

func GetApi(primary string, fallback string) string {
	data, err := Get(primary)
	if err != nil && fallback != "" {
		data, err = Get(fallback)
		if err != nil {
			os.Exit(0)
		}
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return ""
	}
	return strings.TrimSpace(strings.Split(string(data), "\n")[0])
}

func Get(url string) ([]byte, error) {

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("http code error: " + res.Status)
	}

	decoder, err := htmlindex.Get("utf-8")
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(decoder.NewDecoder().Reader(res.Body))

	return buffer.Bytes(), nil
}
