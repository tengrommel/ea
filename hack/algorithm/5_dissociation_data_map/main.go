package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
			result = append(result, right[j])
			j++
		}
	}
	for i < l {
		result = append(result, left[i])
		i++
	}
	for j < r {
		result = append(result, right[j])
		j++
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func makeArray() []int {
	var (
		length = 30
		list   = make([]int, 0)
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)
	return list
}

func quickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	} else {
		var wg sync.WaitGroup // 批量等待
		c := data[0]
		var left, mid, right []int
		mid = append(mid, c) // 第一个数
		for k, v := range data {
			if k == 0 {
				continue
			}
			if c < v {
				right = append(right, v)
			} else if c > v {
				left = append(left, v)
			} else {
				mid = append(mid, c)
			}
		}
		go func() {
			left = quickSort(left)
			wg.Done()
		}()
		go func() {
			right = quickSort(right)
			wg.Done()
		}()
		wg.Add(2)
		wg.Wait()
		res := make([]int, 0)
		if len(left) > 0 {
			res = append(res, left...)
		}
		res = append(res, mid...)
		if len(right) > 0 {
			res = append(res, right...)
		}
		return res
	}
}

func LoadFile() map[string]int {
	myMap := make(map[string]int)
	file, err := os.Open("/where")
	if err != nil {
		fmt.Println("文件打开失败")
		return nil
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		// 计数
		if v, ok := myMap[string(line)]; ok {
			myMap[string(line)] = v + 1
		} else {
			myMap[string(line)] = 1
		}
	}
	return myMap
}

type Pass struct {
	PassWord string
	Times    int
}

func main() {
	myMap := LoadFile()
	fmt.Println(len(myMap)) // 统计数量
	var N int = len(myMap)
	allData := make([]Pass, N, N)
	i := 0
	// 将map迁移到了allData
	for k, v := range myMap {
		allData[i].PassWord = k
		allData[i].Times = v
		i++
	}
	myMap = nil
	runtime.GC()
	debug.FreeOSMemory()

}
