package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/soulteary/simple-kms/internal/define"
	"github.com/soulteary/simple-kms/internal/transformer"
)

// const BASE_API_IN_DOCKER = "http://host.docker.internal"

func GetDockerApiAddr(api string) string {
	return fmt.Sprintf("%s%s%s", define.BASE_API_HOST, define.DEFAULT_PORT, api)
}

func GetAccessKey() string {
	key := os.Getenv("ACCESS_KEY")
	return strings.TrimSpace(key)
}

func Generate() {
	encoded := GetApi(GetDockerApiAddr(define.API_DATA), "")
	seed := GetApi(GetDockerApiAddr(define.API_SEED), "")
	padding := GetApi(GetDockerApiAddr(define.API_PADDING), "")

	accessKey := GetAccessKey()
	secretKey := transformer.Decode(encoded, accessKey, []byte(seed), padding)
	result := fmt.Sprintf("ACCESS_KEY=%s\nSECRET_KEY=%s", accessKey, secretKey)
	fmt.Println(result)
}
