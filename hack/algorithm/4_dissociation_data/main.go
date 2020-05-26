package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := "/some/where"
	passFile, err := os.Open(path)
	defer passFile.Close()
	if err != nil {
		fmt.Println("密码文件打开失败")
		return
	}
	br := bufio.NewReader(passFile)

	saveFile, err := os.Create("some/where")
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, " # ")
		if len(lines) == 3 {
			password := lines[1]
			fmt.Fprintln(save, password)
		}
	}
	save.Flush()
}
