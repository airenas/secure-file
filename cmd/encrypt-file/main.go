package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/airenas/secure/pkg/secure"
)

func main() {
	log.SetOutput(os.Stderr)
	secretPtr := flag.String("s", "", "secret")
	filePtr := flag.String("f", "", "file to encrypt")
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

	log.Printf("Read file " + *filePtr)
	b, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encrypt")

	b, err = secure.Encrypt(b, secret)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)

	log.Print("Done encrypting")
}
