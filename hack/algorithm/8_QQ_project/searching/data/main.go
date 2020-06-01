package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filePath := ""
	file, _ := os.Open(filePath)
	i := 0
	br := bufio.NewReader(file)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if strings.Contains(string(line), "英语") {
			fmt.Println(string(line))
		}
		i++
	}
	fmt.Println(i)
}
