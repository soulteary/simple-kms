package filler

import (
	"math/rand"

	"github.com/google/uuid"
)

func randomBytes(n int) []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func GetFillers() (seed []byte, padding string) {
	return randomBytes(16), string(randomBytes(16))
}

func GetUUID() string {
	return uuid.New().String()
}
