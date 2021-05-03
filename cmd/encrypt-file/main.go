package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/airenas/secure/pkg/secure"
	"github.com/airenas/secure/pkg/util"
	"github.com/pkg/errors"
)

func main() {
	log.SetOutput(os.Stderr)
	printVersion()
	p := util.Params{}
	flag.StringVar(&p.Secret, "s", "", "secret")
	flag.StringVar(&p.File, "f", "", "file to encrypt. File will be encrypted to <file-name.ext.aes>")
	flag.StringVar(&p.FileList, "fl", "", "text file with files to encrypt. One line for one file")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:[params]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if p.Secret == "" {
		p.Secret = os.Getenv("SECRET")
	}
	if p.Secret == "" {
		flag.Usage()
		log.Fatal("No secret")
	}
	if p.FileList == "" && p.File == "" {
		flag.Usage()
		log.Fatal("No file or file list provided")
	}

	if p.FileList != "" && p.File != "" {
		flag.Usage()
		log.Fatal("Only one of <-f> or <-fl> is allowed")
	}

	if p.File != "" {
		err := encryptFile(p.File, p.Secret)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := encryptFiles(p.FileList, p.Secret)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Done")
}

func encryptFile(file, secret string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrapf(err, "can't read '%s'", file)
	}
	b, err = secure.Encrypt(b, secret)
	if err != nil {
		return errors.Wrapf(err, "can't encrypt '%s'", file)
	}
	err = ioutil.WriteFile(file+util.SecureFileExt, b, 0644)
	if err != nil {
		return errors.Wrapf(err, "can't write '%s'", file)
	}
	log.Printf("Encrypted: %s\n", file+util.SecureFileExt)
	return nil
}

func encryptFiles(file, secret string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrapf(err, "can't read '%s'", file)
	}
	lines := strings.Split(string(b), "\n")
	for _, s := range lines {
		if strings.TrimSpace(s) != "" {
			err := encryptFile(s, secret)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

var (
	version string
)

func printVersion() {
	banner := "File encryptor v: %s"
	log.Printf(banner, version)
}
