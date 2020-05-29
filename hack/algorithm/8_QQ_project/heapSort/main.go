package main

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - 1
		HeapSortMax(arr, lastLen)
		if i < length {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}
	return arr
}

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1 // 二叉树深度
		for i := depth; i >= 0; i-- {
			topMax := i
			leftChild := 2*i + 1
			rightChild := 2*i + 2
			if leftChild <= length-1 && arr[leftChild] > arr[topMax] {
				topMax = leftChild
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				topMax = rightChild
			}
			if topMax != i {
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
		return arr
	}
}

func main() {

}
