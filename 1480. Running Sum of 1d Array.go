package main

import "fmt"

func runningSum(nums []int) []int {
	sum := 0
	result := make([]int, len(nums))
	for i, num := range nums {
		fmt.Println(num)
		sum += num
		result[i] = sum
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 1, 1, 1, 1}
	c := []int{3, 1, 2, 10, 1}

	fmt.Println(runningSum(a))
	fmt.Println(runningSum(b))
	fmt.Println(runningSum(c))
}
