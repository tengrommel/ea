package MergeFile

import (
	"bufio"
	"container/list"
	"ea/hack/algorithm/7_distributed_map_reduce/DistributedServer/tools"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Pass struct {
	PassWord string
	times    int
}

func Merge(arr []string, filePath string, dataType int) string {
	myList := list.New()
	for i := 0; i < len(arr); i++ {
		myList.PushBack(arr[i])
	}
	fmt.Println(myList.Len())
	for myList.Len() != 1 {
		e1 := myList.Back()
		myList.Remove(e1)
		e2 := myList.Back()
		myList.Remove(e2)
		if e1 != nil && e2 != nil {
			v1, _ := e1.Value.(string)
			v2, _ := e2.Value.(string)

			fileList := strings.Split(v1, "\\")
			fileName := fileList[len(fileList)-1]
			fileName = strings.Replace(fileName, ".txt", "", -1)
			fileName = fileName[:len(fileName)-1]

			v3 := fileName + getNumberString(v1) + getNumberString(v2) + ".txt"
			fmt.Println(v3)
			Mergre(filePath+v1, filePath+v2, filePath+v3, dataType)
		} else if e1 != nil && e2 == nil {
			v1, _ := e1.Value.(string)
			myList.PushBack(v1)
		} else if e1 == nil && e2 == nil {
			break
		} else {
			break
		}
	}
	return myList.Back().Value.(string)
}

func getNumberString(str1 string) string {
	reg := regexp.MustCompile(`-(\d+).`)
	ss := reg.FindAllStringSubmatch(str1, -1)
	return ss[0][1]
}

func Mergre(path1, path2, pathSave string, dataType int) {
	var isIncreasing func(data1, data2 interface{}) bool
	if dataType == 1 {
		isIncreasing = func(data1, data2 interface{}) bool {
			return data1.(int) < data2.(int)
		}
	} else if dataType == 2 {
		isIncreasing = func(data1, data2 interface{}) bool {
			return data1.(float64) < data2.(float64)
		}
	} else if dataType == 3 {
		isIncreasing = func(data1, data2 interface{}) bool {
			return data1.(string) < data2.(string)
		}
	} else if dataType == 4 {
		isIncreasing = func(data1, data2 interface{}) bool {
			return data1.(Pass).times < data2.(Pass).times
		}
	}

	length1 := tools.GetLineNumber(path1)
	length2 := tools.GetLineNumber(path2)
	fi1, _ := os.Open(path1)
	fi2, _ := os.Open(path2)
	defer fi1.Close()
	defer fi2.Close()
	br1 := bufio.NewReader(fi1)
	br2 := bufio.NewReader(fi2)

	saveFile, _ := os.Create(pathSave)
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)
	i, j := 0, 0
	line1, _, _ := br1.ReadLine()
	line2, _, _ := br2.ReadLine()
	for i < length1 && j < length2 {
		if isIncreasing(line1, line2) {
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++
		} else if isIncreasing(line2, line1) {
			fmt.Fprintln(save, string(line1))
			line1, _, _ = br1.ReadLine()
			i++
		} else {
			fmt.Fprintln(save, string(line1))
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++
			line1, _, _ = br1.ReadLine()
			i++
		}
	}
	for i < length1 {
		fmt.Fprintln(save, string(line1))
		line1, _, _ = br1.ReadLine()
		i++
	}
	for j < length2 {
		fmt.Fprintln(save, string(line2))
		line2, _, _ = br2.ReadLine()
		j++
	}
	save.Flush()
}
