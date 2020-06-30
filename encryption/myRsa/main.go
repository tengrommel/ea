package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 生成rsa的密钥对，并且保存到磁盘文件
func GenerateRsaKey(keySize int) {
	// 1、使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2、通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3、要组织一个pem.Block
	block := pem.Block{
		Type:  "rsa private key", // 这个地方写个字符串就行
		Bytes: derText,
	}
	// 4、pem编码
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
	// 公钥
	// 1、从私钥中取出公钥
	publicKey := privateKey.PublicKey
	// 2、使用x509标准序列化
	derstream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3、将得到的数据放到pem.Block中
	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derstream,
	}
	// 4.pem编码
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

// RSA 加密 公钥加密
func RSAEncrypt(plainText []byte, fileName string) []byte {
	// 打开文件，并且读取文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey, _ := pubInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}
func RSADecrypt(cipherText []byte, fileName string) []byte {
	// 打开文件，并且读取文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	block, _ := pem.Decode(buf)
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}

func main() {
	GenerateRsaKey(1024)
	src := []byte("如果我死了，肯定不是自杀...")
	cipherText := RSAEncrypt(src, "/home/teng/Documents/git/ea/public.pem")
	plainText := RSADecrypt(cipherText, "/home/teng/Documents/git/ea/private.pem")
	fmt.Println(string(plainText))
}
