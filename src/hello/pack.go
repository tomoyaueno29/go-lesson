package main

import (
	"fmt"
	"crypto/hmac"
	"encoding/hex"
	"crypto/sha256"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	
}

func main() {
	const apiKey = "User2Key"
	const apiSecret = "User2Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)

	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)
}