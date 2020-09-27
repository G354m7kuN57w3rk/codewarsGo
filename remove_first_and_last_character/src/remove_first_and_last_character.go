package removefirstandlastcharacter

//RemoveChar remove first and last character of input string
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
