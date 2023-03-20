package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := mapping[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			for _, v2 := range mapping {
				if v2 == t[i] {
					return false
				}
			}
			mapping[s[i]] = t[i]
		}
	}
	return true
}

func main() {
	a := "egg"
	b := "add"
	c := "foo"
	d := "bar"
	e := "paper"
	f := "title"

	fmt.Println(isIsomorphic(a, b))
	fmt.Println(isIsomorphic(c, d))
	fmt.Println(isIsomorphic(e, f))
}
