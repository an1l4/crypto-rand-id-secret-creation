// rand.Int() method
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {
	clientID := generateRandomId()
	clientSecret := generateRandomString(32)
	fmt.Println("client id : ", clientID)
	fmt.Println("*************************************************************")
	fmt.Println("client secret :", clientSecret)

}

// This function is generating cryptographically secure random numbers in a range.
func generateRandomId() *big.Int {
	v, e := rand.Int(rand.Reader, big.NewInt(100000000000))
	if e != nil {
		fmt.Print(e)
	}

	return v
}

func generateRandomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}
