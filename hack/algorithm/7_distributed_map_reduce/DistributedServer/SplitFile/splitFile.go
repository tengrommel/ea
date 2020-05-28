package SplitFile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func evgSplit(num, N int) []int {
	arr := make([]int, 0)
	if num%N == 0 {
		for i := 0; i < N; i++ {
			arr = append(arr, num/N)
		}
	} else {
		evg := (num - num%N) / (N - 1)
		for i := 0; i < N-1; i++ {
			arr = append(arr, evg)
			num -= evg
		}
		arr = append(arr, num)
	}
	return arr
}

func GetLineNumber(filePath string) int {
	fi, err := os.Open(filePath)
	defer fi.Close()
	if err != nil {
		return -1
	}
	i := 0
	br := bufio.NewReader(fi)
	for {
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		i++
	}
	return i
}

func SplitFile(filePath string, savePath string, num int) []string {
	fileList := strings.Split(filePath, "\\")
	fileName := fileList[len(fileList)-1]

	N := GetLineNumber(filePath)
	arrList := evgSplit(N, num)

	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return []string{}
	}

	br := bufio.NewReader(fi)
	tmpFileList := make([]string, 0)
	for i := 0; i < len(arrList); i++ {
		tmpPath := savePath + strings.Replace(fileName, ".txt", "", -1) + strconv.Itoa(i) + ".txt"
		tmpFileList = append(tmpFileList, tmpPath)
		saveFile, _ := os.Create(tmpPath)
		defer saveFile.Close()
		save := bufio.NewWriter(saveFile)
		for j := 0; j < arrList[i]; j++ {
			line, _, _ := br.ReadLine()
			fmt.Fprintln(save, string(line))
		}
		save.Flush()
	}
	fi.Close()
	return tmpFileList
}
