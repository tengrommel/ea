package main

import "fmt"

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
	myFunc       func(data1, data2 interface{}, IsIncreasing bool) bool
}

func (qData *QuickSortData) QuickSort() {
	fmt.Println("原始数据", qData.Data)
	if len(qData.Data) < 20 {
		fmt.Println("插入排序")
		// 插入排序
		BinSearchSort(qData.Data, qData.myFunc, qData.IsIncreasing)
	} else {
		// 快速排序

	}

	fmt.Println("最终数据", qData.Data)
}

func main() {
	type QQ struct {
		password string
		time     int
	}
	mydata := new(QuickSortData)
	mydata.Data = []interface{}{QQ{"addfds", 1}, QQ{"cddfds", 2}, QQ{"bddfds", 2}}
	mydata.IsIncreasing = true
	mydata.myFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
		if isIncreasing {
			return data1.(QQ).time < data2.(QQ).time
		} else {
			return data1.(QQ).time > data2.(QQ).time
		}
	}
	mydata.QuickSort()
	//mydata := new(QuickSortData)
	//mydata.Data = []interface{}{1, 9, 2, 8, 3, 7, 4, 6, 5}
	//mydata.IsIncreasing = true
	//mydata.myFunc = func(data1, data2 interface{}, isIncreasing bool) bool {
	//	if isIncreasing {
	//		return data1.(int) < data2.(int)
	//	} else {
	//		return data1.(int) > data2.(int)
	//	}
	//}
	//mydata.QuickSort()
}
