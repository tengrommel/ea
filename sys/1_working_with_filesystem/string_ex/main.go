package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	b := bytes.NewBuffer(nil)
	b.WriteString("One")
	s1 := b.String()
	b.WriteString("Two")
	s2 := b.String()
	b.Reset()
	b.WriteString("Hey!")
	s3 := b.String()
	fmt.Println("S1: ", s1, "S2:", s2, "S3:", s3)
	stringB := strings.Builder{}
	stringB.WriteString("One")
	fmt.Println(stringB.String())
}
