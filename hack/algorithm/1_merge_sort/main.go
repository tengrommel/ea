package main

import (
	"fmt"
	"math/rand"
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

func main() {
	myArr := makeArray()
	fmt.Println(myArr)
	fmt.Println(MergeSort(myArr))
}
