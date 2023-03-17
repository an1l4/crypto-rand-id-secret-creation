// Prime Method with string output
package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	clientID := generate(128)

	fmt.Println("client Id : ", clientID)

	fmt.Println("*************************************************************")

	clientSecret := generate(124)

	fmt.Println("client Secret : ", clientSecret)
}

func generate(prime int) string {
	RandomCrypto, _ := rand.Prime(rand.Reader, prime)
	state := RandomCrypto.String()
	//fmt.Println(RandomCrypto)
	return state
}
