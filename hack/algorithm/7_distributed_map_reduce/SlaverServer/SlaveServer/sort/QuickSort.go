package sort

import (
	"fmt"
	"math/rand"
	"sync"
)

func FindindexMid(list []interface{}, start int, end int, cur int,
	myFunc func(data1, data2 interface{}, isIncreasing bool) bool, isIncreasing bool) int {
	if start >= end {
		if myFunc(list[start], list[cur], !isIncreasing) {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2
	// 二分查找递归
	if myFunc(list[mid], list[cur], isIncreasing) {
		return FindindexMid(list, start, mid, cur, myFunc, isIncreasing)
	} else {
		return FindindexMid(list, mid+1, end, cur, myFunc, isIncreasing)
	}
}

func BinSearchSort(myList []interface{}, myFunc func(data1, data2 interface{}, isIncreasing bool) bool, isIncreasing bool) []interface{} {
	if len(myList) <= 1 {
		return myList
	} else {
		for i := 1; i < len(myList); i++ {
			p := FindindexMid(myList, 0, i-1, i, myFunc, isIncreasing)
			if p != i {
				for j := i; j > p; j-- {
					myList[j], myList[j-1] = myList[j-1], myList[j]
				}
			}
		}
		return myList
	}
}

type QuickSortData struct {
	Data         []interface{}
	IsIncreasing bool
	MyFunc       func(data1, data2 interface{}, IsIncreasing bool) bool
}

func (qData *QuickSortData) QuickSort() {
	fmt.Println("原始数据", qData.Data)
	if len(qData.Data) < 2 {
		// 插入排序
		BinSearchSort(qData.Data, qData.MyFunc, qData.IsIncreasing)
	} else {
		// 快速排序
		QuickSort(qData.Data, 0, len(qData.Data)-1, qData.MyFunc, qData.IsIncreasing)
	}

	fmt.Println("最终数据", qData.Data)
}

func BinSearchSortIndex(myList []interface{}, start int, end int, myFunc func(data1, data2 interface{},
	isIncreasing bool) bool, isIncreasing bool) []interface{} {
	if end-start <= 1 {
		return myList
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMid(myList, start, i-1, i, myFunc, isIncreasing)
			if p != i {
				for j := i; j > p; j-- {
					myList[j], myList[j-1] = myList[j-1], myList[j]
				}
			}
		}
		return myList
	}
}

func Swap(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSort(arr []interface{}, left, right int,
	myFunc func(data1, data2 interface{}, isIncreasing bool) bool,
	isIncreasing bool) {
	if right-left < 10 {
		BinSearchSortIndex(arr, left, right, myFunc, isIncreasing)
	} else {
		Swap(arr, left, rand.Int()%(right-left)+left)
		vData := arr[left]
		It := left
		gt := right + 1
		i := left + 1
		for i < gt {
			if myFunc(arr[i], vData, isIncreasing) {
				Swap(arr, i, It+1)
				It++
				i++
			} else if myFunc(arr[i], vData, !isIncreasing) {
				Swap(arr, i, gt-1)
				gt--
			} else {
				i++
			}
		}
		Swap(arr, left, It)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			QuickSort(arr, left, It-1, myFunc, isIncreasing)
			wg.Done()
		}()
		go func() {
			QuickSort(arr, gt, right, myFunc, isIncreasing)
			wg.Done()
		}()
		wg.Wait()
	}
}
