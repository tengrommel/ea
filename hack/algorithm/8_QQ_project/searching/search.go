package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	savePath := "path"
	saveFile, err := os.Create(savePath)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)
	dirListx, _ := ioutil.ReadDir("path")
	for j, u := range dirListx {
		fileDir := "path" + u.Name()
		fmt.Println(fileDir, "开始写入")
		fi, err := os.Open(fileDir)
		if err != nil {
			fmt.Println("打开文件失败")
			continue
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			fmt.Fprintln(save, string(line))
		}
		save.Flush()
		fmt.Println(fileDir, "写入成功", j)
	}
	save.Flush()
}
