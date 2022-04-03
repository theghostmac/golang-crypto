package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return privateKey, &privateKey.PublicKey
}

func main() {
	privateKey, publicKey := generateKeyPair(2048)
	fmt.Printf("Private Key: %v\n", privateKey)
	fmt.Printf("Public Key: %v\n", publicKey)
}
