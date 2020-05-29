package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ShellSortGo(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	cpuNum := runtime.NumCPU() // 获取CPU的数量
	wg := sync.WaitGroup{}     // 批量等待
	for gap := len(arr); gap > 0; gap /= 2 {
		wg.Add(cpuNum)
		ch := make(chan int, 10000)
		go func() {
			// 管道写入
			for k := 0; k < gap; k++ {
				ch <- k
			}
			close(ch)
		}()
		for k := 0; k < cpuNum; k++ {
			go func() {
				for v := range ch {
					ShellSortStep(arr, v, gap)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
	fmt.Println(arr)
}

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
