package main

import "fmt"

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap // 跳过步长
		}
		arr[j+gap] = backup
		fmt.Println(arr)
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 2 {
			for i := 0; i < gap; i++ {
				ShellSortStep(arr, i, gap)
			}
			gap /= 2
		}
		return arr
	}
}

func main() {

}
