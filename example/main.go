package main

import (
	"os"
)

var (
	ACCESS_KEY string
	SECRET_KEY string
)

const (
	BASE_API_IN_DOCKER = "http://192.168.88.49"
	BASE_API_HOST      = "http://localhost"
	DEFAULT_PORT       = ":8090"

	API_DATA    = "/config/data"
	API_SEED    = "/config/seed"
	API_PADDING = "/config/padding"
)

func init() {
	encoded := GetApi(GetLocalApiAddr(API_DATA), GetDockerApiAddr(API_DATA))
	seed := GetApi(GetLocalApiAddr(API_SEED), GetDockerApiAddr(API_SEED))
	padding := GetApi(GetLocalApiAddr(API_PADDING), GetDockerApiAddr(API_PADDING))

	accessKey := GetVar(ACCESS_KEY)
	if accessKey == "" {
		accessKey = GetEnv("ACCESS_KEY")
	}

	secretKey := GetVar(SECRET_KEY)
	if secretKey == "" {
		if Encode(GetEnv("SECRET_KEY"), accessKey, seed, padding) != encoded {
			os.Exit(0)
		}
	} else {
		if Encode(secretKey, accessKey, seed, padding) != encoded {
			os.Exit(0)
		}
	}
}
