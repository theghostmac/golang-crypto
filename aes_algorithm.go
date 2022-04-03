package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

// generateKey
func generateKey() []byte {
	randomByte := make([]byte, 32)
	numBytesRead, err := rand.Read(randomByte)
	if err != nil {
		log.Fatal("Error generating random key.", err)
	}
	if numBytesRead != 32 {
		log.Fatal("Error generating 32 random bytes for key.")
	}
	return randomByte
}

func encrypt(key, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]bytem aes.BlockSize+len(message))

	initVectorNonce := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initVectorNonce)
	if err != nil {
		return nil, err
	}

	cipherFeedBack := cipher.NewCFBEncrypter(block, initVectorNonce)
	cipherFeedBack.XORKeyStream(cipherText[aes.BlockSize:], message)

	return cipherText, nil
}

func decrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	initVectorNonce := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize]

	cipher := cipher.NewCFBDecrypter(block, initVectorNonce)
	cipher.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func main() {
	keyFileDate, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal("Unable to read key file contents.", err)
	}

	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Unable to read key file contents.", err)
	}

	if decryptFlag {
		message, err := decrypt(keyFileDate, fileData)
		if err != nil {
			log.Fatal("Error decrypting.", err)
		}
		fmt.Printf("%s", message)
	} else {
		cipherText, err := encrypt(keyFileDate, fileData)
		if err != nil {
			log.Fatal("Error encrypting. ", err)
		}
		fmt.Printf("%s", cipherText)
	}
}