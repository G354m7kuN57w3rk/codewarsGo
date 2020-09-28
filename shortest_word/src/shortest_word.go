package shortestword

import "strings"

//FindShort returns the length of the shortest word of input string
func FindShort(s string) int {
	strSplit := strings.Split(s, " ")
	ans := len(strSplit[0])
	for _, v := range strSplit {
		if ans > len(v) {
			ans = len(v)
		}
	}
	return ans
}
