package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	password := "123456"
	h1 := md5.New()
	h1.Write([]byte(password))
	fmt.Println(hex.EncodeToString(h1.Sum(nil)))

	password1 := "123456"
	h2 := md5.New()
	h2.Write([]byte(password1))
	fmt.Println(hex.EncodeToString(h2.Sum(nil)))
}
