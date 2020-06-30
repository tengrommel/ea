package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	src := []byte("GenerateHmac(plainText, key []byte) []byte")
	key := []byte("hello world")
	hamc1 := GenerateHmac(src, key)
	b1 := VerifyHmac(src, key, hamc1)
	fmt.Printf("校验结果： %t\n", b1)
}

// 生消息认证码
func GenerateHmac(plainText, key []byte) []byte {
	// 1、创建一个哈希接口，需要指定使用的哈希算法，和秘钥
	myHash := hmac.New(sha1.New, key)
	// 2、给哈希对象添加数据
	myHash.Write(plainText)
	// 3、计算散列值
	hashText := myHash.Sum(nil)
	return hashText
}

// 校验消息认证码
func VerifyHmac(plainText, key, hashText []byte) bool {
	// 1、创建哈希接口，需要指定使用的哈希算法，和秘钥
	myHash := hmac.New(sha1.New, key)
	// 2、给哈希对象添加数据
	myHash.Write(plainText)
	// 3、计算散列值
	hamc1 := myHash.Sum(nil)
	// 4、两个散列值比较
	return hmac.Equal(hashText, hamc1)
}
