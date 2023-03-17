// rand.Read() method
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	clientID := generateRandomString(16)
	clientSecret := generateRandomString(32)

	fmt.Printf("Client ID: %s\n", clientID)
	fmt.Println("**************************************************************")
	fmt.Printf("Client Secret: %s\n", clientSecret)
}

func generateRandomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}
