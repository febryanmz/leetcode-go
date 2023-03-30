package main

import "fmt"

func main() {

	//Example 1:
	//Input:
	haystack := "sadbutsad"
	needle := "sad"

	fmt.Println(haystack, needle)
	fmt.Println(strStr(haystack, needle))

	//Output: 0
	//Explanation: "sad" occurs at index 0 and 6.
	//The first occurrence is at index 0, so we return 0.
	//

	//Example 2:

	//Input:
	haystack2 := "leetcode"
	needle2 := "leeto"

	fmt.Println(haystack2, needle2)
	fmt.Println(strStr(haystack2, needle2))

	//Output: -1
	//Explanation: "leeto" did not occur in "leetcode", so we return -1.

	haystack3 := "dimanaakudimana"
	needle3 := "aku"
	fmt.Println(haystack3, needle3)
	fmt.Println(strStr(haystack3, needle3))

}

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	if m == 0 {
		return 0
	}

	for i := 0; i <= n-m; i++ {
		fmt.Println("i : ", i, "haystack : ", haystack[i:i+m])
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}
