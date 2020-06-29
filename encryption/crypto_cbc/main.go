package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 1、求出最后一个组中剩余的字节数 28 % 8 3 .. 4
	padNum := blockSize - len(plainText)%blockSize
	// 2、创建一个新的切片，长度==padNum 每个字节值 byte(padNum)
	char := []byte{byte(padNum)} // 长度为1
	// 切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	// 3、newPlain数组追加到原始明文的后面
	return append(plainText, newPlain...)
}

func unPaddingLastGroup(plainText []byte) []byte {
	// 1、拿去切片中的最后一个字节
	lenght := len(plainText)
	lastChar := plainText[lenght-1] // byte类型
	number := int(lastChar)         // 尾部填充的字节数
	return plainText[:lenght-number]
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
	cipher.NewCBCEncrypter(block, iv)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4、加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)
	return cipherText
}

// des的CBC加密
func main() {

}
