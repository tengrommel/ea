package logn

import (
	"math/rand"
	"sort"
	"time"
)

var testSet = generateIntSlice(100000)
var randSearchNum = randNumber(testSet)

func generateIntSlice(n int) []int {
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x = append(x, rand.Intn(10000))
	}
	sort.Ints(x)
	return x
}

func randNumber(testSet []int) int {
	rand.Seed(time.Now().Unix())
	randSeed := rand.Int() % len(testSet)
	randSearchNum := testSet[randSeed]
	return randSearchNum
}

func binarySearch(intSlice []int, searchNum int) {
	sort.SearchInts(intSlice, searchNum)
}

func binarySearchTimings(n int) {
	binarySearch(testSet[0:n], randSearchNum)
}
