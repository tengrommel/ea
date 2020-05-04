package main

import (
	"fmt"
)

/**
Given a string, find the length of the longest substringwithout repeating characters.
通过slice简化map操作，利用快慢指针解决问题
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result := 0
	left, right := 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(result int, i int) int {
	if result > i {
		return result
	}
	return i
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dddfa"))
}
