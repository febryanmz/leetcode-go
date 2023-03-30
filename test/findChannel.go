package main

import "fmt"

func solve(n int, k int) int {
	if n > 1 {
		return (solve(n-1, k)+k-1)%n + 1
	}
	return 1
}

func solve2(n int, k int) int {
	// n -> nos of channels; k -> nos of channels to skip
	channel := make([]bool, n)
	// ith element is false if ith channel is deleted
	for i := range channel {
		channel[i] = true
	}
	i, cnt := 0, n // i -> current channel; cnt -> nos of undeleted channels
	for cnt > 1 {
		for j := 0; j < k; j++ {
			for !channel[i%n] {
				fmt.Println("ini dalem", channel)
				i++
			}
			fmt.Println("chann", channel)
			i++
		}
		channel[(i-1)%n] = false
		cnt--
	}
	fmt.Println("3 :", channel)
	for i = 0; !channel[i]; i++ {
		fmt.Println("4 :", channel)
	}
	fmt.Println("out", channel)
	return i + 1
}

func main() {
	fmt.Println(solve(5, 2))
	fmt.Println(solve2(5, 2))
}
