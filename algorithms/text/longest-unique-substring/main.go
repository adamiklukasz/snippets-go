package main

import (
	"fmt"
)

// a = 0x61

func longestUniqueSubstring(input string) string {
	lastIndexForChar := make([]int, 255)
	for i := 0; i < len(lastIndexForChar); i++ {
		lastIndexForChar[i] = -1
	}

	lastMinIndex := 0
	longestSubstring := ""

	for i := 0; i < len(input); i++ {
		c := input[i]
		if lastIndexForChar[c] == -1 {
			lastIndexForChar[c] = i
			if len(input[lastMinIndex:i+1]) > len(longestSubstring) {
				longestSubstring = input[lastMinIndex : i+1]
			}
		} else {
			lastMinIndex = lastIndexForChar[c] + 1
			lastIndexForChar[c] = i
		}
	}

	return longestSubstring
}

func main() {
	res := longestUniqueSubstring("abcadcefac")
	fmt.Printf("res=%#v\n", res)
}
