package main

import "bytes"

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

// des的CBC加密
func main() {

}
