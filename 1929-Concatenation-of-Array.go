package main

import "fmt"

// --Solution 1--
// func getConcatenation(nums []int) []int {
// 	return append(nums, nums...)
// }

// --Solution 2--
// func getConcatenation(nums []int) []int {
// 	var numsLen = len(nums)
// 	var result = make([]int, numsLen*2)
// 	for i := 0; i < numsLen; i++ {
// 		result[i] = nums[i]
// 		result[numsLen+i] = nums[i]
// 	}
// 	return result
// }

// --Solution 3--
func getConcatenation(nums []int) []int {
	var out = make([]int, len(nums)*2)
	copy(out, nums)

	for index, value := range nums {
		out[index+len(nums)] = value
	}

	return out
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3, 4, 5}))
	fmt.Println(getConcatenation([]int{6, 7, 8, 9, 10, 11}))
}
