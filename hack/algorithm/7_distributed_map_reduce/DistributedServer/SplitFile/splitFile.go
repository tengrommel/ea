package SplitFile

import (
	"bufio"
	"ea/hack/algorithm/7_distributed_map_reduce/DistributedServer/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitFile(filePath string, savePath string, num int) []string {
	fileList := strings.Split(filePath, "\\")
	fileName := fileList[len(fileList)-1]
	N := tools.GetLineNumber(filePath)
	arrList := tools.EvgSplit(N, num)
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
