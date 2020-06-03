package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 制造索引
func makeIndex() []int {
	N := 156774123
	indexData := make([]int, N, N)
	path := "path"
	file, _ := os.Open(path)
	br := bufio.NewReader(file)
	indexData[0] = 0
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lineString := string(line)
		indexData[i] = len(lineString)
		if i < N {
			indexData[i] = len(lineString)
		}
		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}
	}
	file.Close()
	fmt.Println("第一步读取长度索引", i)
	for j := 0; j < len(indexData)-1; j++ {
		indexData[j+1] += indexData[j]
	}
	fmt.Println("第二部数组索引叠加完成")
	return indexData
}

func main() {
	path := ""
	file, _ := os.Open(path)
	file.Seek(0, 0) // 移动到文件开头
	lengthList := makeIndex()
	for {
		var lineNum int
		fmt.Scanf("%d", &lineNum)
		file.Seek(int64(lengthList[lineNum]), 0) // 调到指定行
		b := make([]byte, 12+4+16, 12+4+16)
		length, _ := file.Read(b)
		var endPos int
		for i := 0; i < length-1; i++ {
			if b[i] == '\n' && i > 5+4+6 {
				endPos = i
				break
			}
		}
		fmt.Println("显示数据", string(b[:endPos]))
	}
}
