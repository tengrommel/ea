package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func SignatureRSA(plainText []byte, fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	myHash := sha512.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(buf)
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

func VerifyRSA(plainText, sigText []byte, pubFileName string) bool {
	file, err := os.Open(pubFileName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	hashText := sha512.Sum512(plainText)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], sigText)
	if err == nil {
		return true
	}
	return false
}

func main() {
	src := []byte("的风景好大风和")
	sigText := SignatureRSA(src, "/home/teng/Documents/git/ea/private.pem")
	b1 := VerifyRSA(src, sigText, "public.pem")
	fmt.Println(b1)
}
