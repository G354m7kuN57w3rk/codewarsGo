package twotoone

import (
	"sort"
	"strings"
)

//TwoToOne concat two strings and returns a sorted string containing distinct characters
func TwoToOne(s1 string, s2 string) string {
	var ans string
	strSplit := strings.Split(s1+s2, "")
	sort.Strings(strSplit)
	for _, v := range strSplit {
		chr := string(v)
		if !strings.Contains(ans, chr) {
			ans += chr
		}
	}
	return ans
}
