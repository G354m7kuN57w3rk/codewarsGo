package mumbling

import (
	"strings"
)

//Accum returns a mumbling like "abcd" -> "A-Bb-Ccc-Dddd"
func Accum(s string) string {
	var ans string
	for i, v := range s {
		ans += strings.ToUpper(string(v)) + strings.Repeat(strings.ToLower(string(v)), i)
		if i < len(s)-1 {
			ans += "-"
		}
	}
	return ans
}
