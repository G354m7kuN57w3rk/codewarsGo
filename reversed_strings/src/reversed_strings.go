package reversestrings

//Solution reverse input string
func Solution(word string) (ans string) {
	for _, v := range word {
		ans = string(v) + ans
	}
	return
}
