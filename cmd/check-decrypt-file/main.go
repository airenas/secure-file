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

	p := util.Params{}
	flag.StringVar(&p.Secret, "s", "", "secret")
	flag.StringVar(&p.File, "f", "", "result file, will look after <file.ext>.aes for decryption")
	flag.StringVar(&p.FileList, "fl", "", "text file with files to decrypt. One line for one file")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:[params] [output-file to stdout]\n", os.Args[0])
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
	if p.FileList == "" || p.File == "" {
		flag.Usage()
		log.Fatal("No file or file list provided")
	}

	if p.FileList != "" && p.File != "" {
		flag.Usage()
		log.Fatal("Only one of <-f> or <-fl> is allowed")
	}

	if p.File != "" {
		err := decryptFile(p.File, p.Secret)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := decryptFiles(p.FileList, p.Secret)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Done")
}

func decryptFile(file, secret string) error {
	if fileUpToDate(file) {
		log.Printf("File exists: %s", file)
		return nil
	}
	b, err := ioutil.ReadFile(file + util.SecureFileExt)
	if err != nil {
		return errors.Wrapf(err, "can't read '%s'", file+util.SecureFileExt)
	}
	b, err = secure.Decrypt(b, secret)
	if err != nil {
		return errors.Wrapf(err, "can't decrypt '%s'", file+util.SecureFileExt)
	}
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return errors.Wrapf(err, "can't write '%s'", file)
	}

	log.Printf("Decrypted: %s\n", file)
	return nil
}

func decryptFiles(file, secret string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrapf(err, "can't read '%s'", file)
	}
	lines := strings.Split(string(b), "\n")
	for _, s := range lines {
		if strings.TrimSpace(s) != "" {
			err := decryptFile(s, secret)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func fileUpToDate(filename string) bool {
	outFile, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	inFile, err := os.Stat(filename + util.SecureFileExt)
	if os.IsNotExist(err) {
		return false
	}
	return !outFile.IsDir() && !inFile.IsDir() && outFile.ModTime().After(inFile.ModTime())
}
