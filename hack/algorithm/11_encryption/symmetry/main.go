package main

import (
	"bytes"
	"fmt"
)

// 追加大小
func PKCSPz(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize // 补充的位数
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padText...)
}

// 去掉追加的数据
func PKCSUnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:unPadding]
}

func main() {
	key := []byte("abcdefg") // 密码
	//data :=[]byte("锄禾日当午")
	fmt.Println(string(PKCSPz(key, 4)))
	fmt.Println(len(string(PKCSPz(key, 4))))
	fmt.Println(len(string(PKCSUnPadding(key))))
	fmt.Println(string(PKCSUnPadding(key)))
}
