package main

import (
	sha2562 "crypto/sha256"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 30*1024)
	sha256 := sha2562.New()
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			_, err := sha256.Write(buffer[:n])
			if err != nil {
				log.Fatal()
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum := sha256.Sum(nil)
	log.Printf("%x\n", sum)
}
