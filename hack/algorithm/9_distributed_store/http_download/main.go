package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://192.168.1.38:8080/group1/default/14/32/6/1.html"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("1.html")
	if err != nil {
		panic(err)
	}
	io.Copy(file, res.Body)
}
