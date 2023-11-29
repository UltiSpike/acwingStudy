package main

import (
	"fmt"
)

func removeDuplicates(s string) string {
	var ans []byte
	var n = -1
	for n != len(ans) {
		n = len(ans)
		ans = remove(s)
		s = string(ans)

	}

	return string(ans)
}

func remove(s string) []byte {
	var ans []byte
	s = s + " "
	for i := 0; i < len(s)-1; i++ {

		if s[i] == s[i+1] {
			i += 2
		} else {
			ans = append(ans, s[i])
		}
		fmt.Println(string(ans))
	}

	return ans
}
