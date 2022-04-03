package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
import (
	crypto/sha256
)

func UsageSample()  {
	fmt.Println("Usage: ", os.Args[0], " <filepath>")
	fmt.Println("Example: ", os.Args[0], " file.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		UsageSample()
		os.Exit(1)
	}
	return os.Args[1]
}

func main()  {
	filename := checkArgs()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}