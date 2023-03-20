package main

import "fmt"

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func isSubsequence(s string, t string) bool {
	var status bool
	for _, dariS := range s {
		for _, dariT := range t {
			if dariS == dariT {
				status = true
			}
		}
	}
	return status

}

func isSubsequence2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sIndex, tIndex := 0, 0
	for tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
			if sIndex == len(s) {
				return true
			}
		}
		tIndex++
	}
	return false
}

func main() {
	a := "abc"
	b := "ahbgdc"
	c := "axc"
	d := "ahbgdc"

	fmt.Println(isSubsequence(a, b))
	fmt.Println(isSubsequence(c, d))
	fmt.Println(isSubsequence2(a, b))
	fmt.Println(isSubsequence2(c, d))
}
