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

func runSum(nums []int) []int {
	sums := 0
	newArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sums = sums + nums[i]
		newArr[i] = sums
	}
	return newArr
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 1, 1, 1, 1}
	c := []int{3, 1, 2, 10, 1}

	fmt.Println(runningSum(a))
	fmt.Println(runningSum(b))
	fmt.Println(runningSum(c))

	fmt.Println("test kedua :", runSum(a))
	fmt.Println("test kedua :", runSum(b))
	fmt.Println("test kedua :", runSum(c))
}
