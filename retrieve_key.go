package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func exportPEMStringToPrivateKey(privatekey []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(privatekey)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func exportPEMStringToPublicKey(publickey []byte) *rsa.PublicKey {
	block, _ := pem.Decode(publickey)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}

func readKeyFromFile(filename string) []byte {

}