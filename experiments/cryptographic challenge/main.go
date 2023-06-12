package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateChallenge(key []byte) (message, mac []byte) {
	message = make([]byte, 16)
	_, err := rand.Read(message)
	if err != nil {
		panic(err)
	}

	mac = hmac.New(sha256.New, key).Sum(message)

	return
}

func SolveChallenge(message, mac, key []byte) bool {
	expected := hmac.New(sha256.New, key).Sum(message)

	return hmac.Equal(mac, expected)
}

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	message, mac := GenerateChallenge(key)

	messageB64 := base64.StdEncoding.EncodeToString(message)
	macB64 := base64.StdEncoding.EncodeToString(mac)

	fmt.Println("Challenge:", messageB64)
	fmt.Println("MAC:", macB64)

	solved := SolveChallenge(message, mac, key)

	fmt.Println("Solved:", solved)
}
