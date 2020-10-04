package duplicateencoder

import "strings"

//DuplicateEncode returns a new encoded string
func DuplicateEncode(word string) string {
	var count int
	var ans string
	word = strings.ToLower(word)
	for _, v := range word {
		for _, x := range word {
			if v == x {
				count++
			}
		}
		if count > 1 {
			ans += ")"
		} else {
			ans += "("
		}
		count = 0
	}
	return ans
}
