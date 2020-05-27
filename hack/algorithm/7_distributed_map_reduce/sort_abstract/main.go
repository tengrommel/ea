package main

//func FindindexMid(list []interface{}, start int, end int, cur int, myFunc func (data1, data2 interface{}) bool) int {
//	if start >= end {
//		if myFunc(list[start], list[cur]) {
//			return cur
//		} else {
//			return start
//		}
//	}
//	mid := (start + end)/2
//	// 二分查找递归
//	if myFunc(list[mid], list[cur]) {
//		return FindindexMid(list, start, mid, cur, myFunc)
//	} else {
//		return FindindexMid(list, mid+1, end, cur, myFunc)
//	}
//}
//
//func BinSearchSort(myList []interface{}, myFunc func (data1, data2 interface{}) bool) []interface{} {
//	if len(myList) <= 1 {
//		return myList
//	} else {
//		for i:=1;i<len(myList);i++ {
//			p := FindindexMid(myList, 0, i-1, i, myFunc)
//			if p != i {
//				for j:=i;j>p;j-- {
//					myList[j], myList[j-1] = myList[j-1], myList[j]
//				}
//			}
//		}
//		return myList
//	}
//}
//
//func BinSearchSortIndex(myList []int, start int, end int) []int {
//	if end-start <= 1 {
//		return myList
//	}else {
//		for i:=start+1;i<=end;i++ {
//			p := FindindexMid(myList, start, i-1, i)
//			if p!=i {
//				for j:=i;j>p;j-- {
//					myList[j], myList[j-1]=myList[j-1],myList[j]
//				}
//			}
//		}
//		return myList
//	}
//}
//
//func QuickSortCall(arr []int) []int {
//	if len(arr) < 10 {
//		return BinSearchSort(arr)
//	} else {
//		QuickSort(arr, 0, len(arr)-1)
//		return arr
//	}
//}
//
//func Swap(arr []int, i, j int)  {
//	arr[i], arr[j] = arr[j], arr[i]
//}
//
//func QuickSort(arr []int, left, right int)  {
//	if right-left < 10 {
//		BinSearchSortIndex(arr, left, right)
//	} else {
//		Swap(arr, left, rand.Int()%(right-left)+left)
//		vData := arr[left]
//		It := left
//		gt := right + 1
//		i := left + 1
//		for i<gt {
//			if arr[i] > vData{
//				Swap(arr, i, It+1)
//				It++
//				i++
//			} else if arr[i] < vData {
//				Swap(arr, i, gt-1)
//				gt--
//			} else {
//				i++
//			}
//		}
//		Swap(arr, left, It)
//		var wg sync.WaitGroup
//		wg.Add(2)
//		go func() {
//			QuickSort(arr, left, It-1)
//			wg.Done()
//		}()
//		go func() {
//			QuickSort(arr, gt, right)
//			wg.Done()
//		}()
//		wg.Wait()
//	}
//}
