package main

import "fmt"

func pivotIndex(nums []int) int {

	sumL := 0
	sumR := 0

	for _, n := range nums {
		sumR += n
	}

	for i, n := range nums {
		sumR -= n
		if sumL == sumR {
			return i
		}
		sumL += n
	}

	return -1

}

func main() {
	a := []int{1, 7, 3, 6, 5, 6}
	b := []int{1, 2, 3}
	c := []int{2, 1, -1}

	runA := pivotIndex(a)
	runB := pivotIndex(b)
	runC := pivotIndex(c)

	fmt.Println(runA)
	fmt.Println(runB)
	fmt.Println(runC)
}
