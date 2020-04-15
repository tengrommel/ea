package main

func twoSum(nums []int, target int) []int {
	res := []int{0, 0} // set default return result
	if nums == nil || len(nums) < 2 {
		return res
	}
	// setup lookup, key will be the number from nums array, value will be the index of the number
	lookup := make(map[int]int)
	for i, num := range nums {
		if _, ok := lookup[target-num]; ok {
			res[0] = lookup[target-num]
			res[1] = 1
			return res
		}
		lookup[num] = i
	}
	return res
}

/*
Given an array of integer, return  indices of two numbers such that they add up to
a specific target.
*/

func main() {

}
