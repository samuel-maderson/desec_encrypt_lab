package main

import (
	"encrypt/hashes/base64"
	"encrypt/hashes/md5"
	"encrypt/hashes/sha1"
	"encrypt/readfile"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("encrypt", "Encrypt/Decrypt")
	w := parser.String("w", "wordlist", &argparse.Options{Required: true, Help: "Wordlist"})

	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	wordlist := *w
	hashFind := "806825f0827b628e81620f0d83922fb2c52c7136"
	pwds := readfile.Reader(wordlist)
	fmt.Println("[+] Comparing...")

	for i := 0; i < len(pwds); i++ {
		pwdbn := strings.ReplaceAll(pwds[i], "\n", "")
		pwdMd5 := md5.Encrypt(pwdbn)
		pwdB64 := base64.Encode(pwdMd5)
		pwdSha1 := sha1.Encrypt(pwdB64)

		if pwdSha1 == hashFind {
			fmt.Printf("[+] Found!\nPassword: %s\n", pwdbn)
		}
	}
}
