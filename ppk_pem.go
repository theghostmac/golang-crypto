package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func exportPubKeyAsPEMString(publickey *rsa.PublicKey) string {
	publickeyPEM := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(publickey),
		},
	))
	return publickeyPEM
}

func exportPrivKeyAsPEMString(privatekey *rsa.PrivateKey) string {
	privatekeyPEM := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
		},
	))
	return privatekeyPEM
}

func saveKeyToAFile(keyPEM, filename string) {
	pemBytes := []byte(keyPEM)
	ioutil.WriteFile(filename, pemBytes, 0400)
}

func main() {
	privateKey, publicKey := generateKeyPair(2048)
	fmt.Printf("Private key: %v\n", privateKey)
	fmt.Printf("Public key: %v\n", publicKey)

	privateKeyString := exportPrivKeyAsPEMString(privateKey)
	publicKeyString := exportPubKeyAsPEMString(publicKey)

	fmt.Println(privateKeyString)
	fmt.Println(publicKeyString)

	saveKeyToAFile(privateKeyString, "privatekey.pem")
	saveKeyToAFile(publicKeyString, "publickey.pem")
}