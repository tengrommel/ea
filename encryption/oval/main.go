package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 生成密钥对
func GenerateEccKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		panic(err)
	}
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	block := pem.Block{
		Type:  "ecdsa private key",
		Bytes: derText,
	}
	file, err := os.Create("eccPrivate.pem")
	pem.Encode(file, &block)
	file.Close()
	publicKey := privateKey.PublicKey
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	block = pem.Block{
		Type:  "ecdsa public key",
		Bytes: derText,
	}
	file, err = os.Create("eccPublic.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

func main() {

}
