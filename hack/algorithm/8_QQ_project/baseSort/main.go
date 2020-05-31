package main

import "fmt"

func main() {
	var arr [3][]int
	myArr := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 3}
	for i := 0; i < len(myArr); i++ {
		arr[myArr[i]-1] = append(arr[myArr[i]-1], myArr[i])
	}
	fmt.Println(arr)
}
