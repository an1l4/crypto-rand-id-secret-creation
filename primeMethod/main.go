//Prime Method with big int output
package main

import (
	"crypto/rand"
	"fmt"
)

// it will generate different output each time.
func main() {
	RandomCrypto, _ := rand.Prime(rand.Reader, 128)
	fmt.Println(RandomCrypto)

}
