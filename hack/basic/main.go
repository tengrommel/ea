package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	// Print HTTP Status
	fmt.Println(resp.Status)
	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
	r2, err := http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	r2.Body.Close()
	fmt.Println(resp.Status)
	form := url.Values{}
	form.Add("foo", "bar")
	r3, err := http.Post(
		"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()))
	defer r3.Body.Close()
}
