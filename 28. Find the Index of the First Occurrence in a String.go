package main

import "fmt"

func main() {

	//Example 1:
	//Input:
	haystack := "sadbutsad"
	needle := "sad"

	fmt.Println(haystack, needle)

	//Output: 0
	//Explanation: "sad" occurs at index 0 and 6.
	//The first occurrence is at index 0, so we return 0.
	//

	//Example 2:

	//Input:
	haystack2 := "leetcode"
	needle2 := "leeto"

	fmt.Println(haystack2, needle2)

	//Output: -1
	//Explanation: "leeto" did not occur in "leetcode", so we return -1.

}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return len(needle) - len(haystack)
	}
	for i := 0; i < len(haystack); i++ {
		if
	}
}
