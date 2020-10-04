package stringendswith

import "strings"

//solution returns true if ending is contained in str
func solution(str, ending string) bool {
	return strings.Contains(str, ending)
}
