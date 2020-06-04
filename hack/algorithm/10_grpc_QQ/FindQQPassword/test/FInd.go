package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BinSearchDisk(file *os.File, fileIndex *os.File, QQ int) string {
	const N = 156774124
	low := 0
	high := N - 1
	for low < high {
		mid := (low + high) / 2
		fileIndex.Seek(int64(mid*4), 0)
		bx := make([]byte, 4, 4)
		_, _ = fileIndex.Read(bx)
		var depData uint32 = uint32(binary.BigEndian.Uint32(bx))
		file.Seek(int64(depData), 0)
		b := make([]byte, 12+4+16, 12+4+16)
		length, _ := file.Read(b)
		var endPos int
		for i := 0; i < length-1; i++ {
			if b[i] == '\n' && i >= 5+4+6 {
				endPos = i
				break
			}
		}
		midString := string(b[:endPos])
		fmt.Println("mid 数据", midString)
		midList := strings.Split(midString, "-------")
		midQQ, _ := strconv.Atoi(midList[0])
		if midQQ > QQ {
			high = mid - 1
		} else if midQQ < QQ {
			low = mid + 1
		} else {
			return midString
		}
	}
	return "没有找到"
}

func main() {

}
