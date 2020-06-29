package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 1、求出最后一个组中剩余的字节数 28 % 8 3 .. 4
	padNum := blockSize - len(plainText)%blockSize
	// 2、创建一个新的切片，长度==padNum 每个字节值 byte(padNum)
	char := []byte{byte(padNum)} // 长度为1
	// 切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	// 3、newPlain数组追加到原始明文的后面
	newText := append(plainText, newPlain...)
	return newText
}

func unPaddingLastGroup(plainText []byte) []byte {
	// 1、拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1] // byte类型
	number := int(lastChar)         // 尾部填充的字节数
	return plainText[:length-number]
}

// des 加密
func desEncrypt(plainText, key []byte) []byte {
	// 1、建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2、明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 3、创建一个使用cbc分组接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4、加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)
	return cipherText
}

// des的解密
func desDecrypt(cipherText, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2、创建一个使用cbc解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3、解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 4、cipherText 存储的是明文
	plainText := unPaddingLastGroup(cipherText)
	return plainText
}

func main() {
	fmt.Println("Des 加密 解密")
	key := []byte("01234567")
	src := []byte("戀愛靠機會，分手靠智慧。曹操歷經滄桑，心裡想的大概是：好頭不如好尾。")
	cipherText := desEncrypt(src, key)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("解密之后的数据：%s\n", string(plainText))
}
