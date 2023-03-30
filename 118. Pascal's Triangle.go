package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res

}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(5))
	fmt.Println(generate(10))
}
