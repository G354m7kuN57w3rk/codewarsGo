package vowelcount

//GetCount returns number of vowels from input string
func GetCount(str string) (count int) {
	for _, v := range str {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
