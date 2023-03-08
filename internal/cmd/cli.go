package cmd

import (
	"os"
	"strings"
)

func GetAccessKey() string {
	key := os.Getenv("ACCESS_KEY")
	return strings.TrimSpace(key)
}
