package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func GetFileMd5(filepath string) string {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return ""
	}
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return ""
	}
	md5Hash.Sum(nil)
	return fmt.Sprintf("%x", md5Hash.Sum(nil))
}

func main() {

}
