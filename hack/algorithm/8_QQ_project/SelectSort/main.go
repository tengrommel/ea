package main

import (
	"fmt"
)

// 返回一个最大值
func SelectMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0] // 假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
		return arr
	}
}

// topk
// 百度、淘宝搜索、---日志
// 关键字搜索次数最多 topk100
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5} // 数组排序
	fmt.Println(arr)
}
