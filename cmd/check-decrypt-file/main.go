package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/airenas/secure/pkg/secure"
)

const secureExt = ".aes"

func main() {
	log.SetOutput(os.Stderr)
	secretPtr := flag.String("s", "", "secret")
	filePtr := flag.String("f", "", "result file, will look after <file.ext>.aes for decryption")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:[params] [output-file to stdout]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *filePtr == "" {
		flag.Usage()
		log.Fatal("No file")
	}
	secret := *secretPtr
	if secret == "" {
		secret = os.Getenv("SECRET")
	}
	if secret == "" {
		flag.Usage()
		log.Fatal("No secret")
	}

	if fileUpToDate(*filePtr) {
		log.Printf("File exists")
		return
	}

	log.Printf("Read file " + *filePtr + secureExt)
	b, err := ioutil.ReadFile(*filePtr + secureExt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Decrypt")

	b, err = secure.Decrypt(b, secret)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(*filePtr, b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done decrypting")
}

func fileUpToDate(filename string) bool {
	outFile, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	inFile, err := os.Stat(filename + secureExt)
	if os.IsNotExist(err) {
		return false
	}
	return !outFile.IsDir() && !inFile.IsDir() && outFile.ModTime().After(inFile.ModTime())
}
