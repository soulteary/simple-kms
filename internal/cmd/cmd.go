package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/soulteary/simple-kms/internal/define"
	"github.com/soulteary/simple-kms/internal/filler"
	"github.com/soulteary/simple-kms/internal/transformer"
)

// const BASE_API_IN_DOCKER = "http://host.docker.internal"

func GetDockerApiAddr(api string) string {
	return fmt.Sprintf("%s%s%s", define.BASE_API_HOST, define.DEFAULT_PORT, api)
}

func GetAccessKey(isCli bool) string {
	key := strings.TrimSpace(os.Getenv("ACCESS_KEY"))
	if key == "" && !isCli {
		key = strings.TrimSpace(filler.GetUUID())
		fmt.Println(fmt.Sprintf("ACCESS_KEY is empty, filling with random uuid: %s", key))
	}
	return key
}

func GenerateSecretByAPI() {
	encoded := GetApi(GetDockerApiAddr(define.API_DATA), "")
	seed := GetApi(GetDockerApiAddr(define.API_SEED), "")
	padding := GetApi(GetDockerApiAddr(define.API_PADDING), "")

	accessKey := GetAccessKey(true)
	if accessKey == "" {
		fmt.Println("ACCESS_KEY is empty")
		return
	}

	secretKey := transformer.Decode(encoded, accessKey, []byte(seed), padding)
	result := fmt.Sprintf("ACCESS_KEY=%s\nSECRET_KEY=%s", accessKey, secretKey)
	fmt.Println(result)
}

func GenerateSecretByCli() {
	result := fmt.Sprintf("ACCESS_KEY=%s\nSECRET_KEY=%s", filler.GetUUID(), filler.GetUUID())
	fmt.Println(result)
}
