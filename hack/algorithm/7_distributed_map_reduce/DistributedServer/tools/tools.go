package tools

import (
	"bufio"
	"io"
	"os"
)

func EvgSplit(num, N int) []int {
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
