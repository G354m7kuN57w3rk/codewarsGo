package printererrors

import (
	"fmt"
	"strings"
)

//PrinterError returns error rate of input string
func PrinterError(s string) string {
	col := "abcdefghijklm"
	var ans int
	strSplit := strings.Split(s, "")
	for _, v := range strSplit {
		chr := string(v)
		if !strings.Contains(col, chr) {
			ans++
		}
	}
	return fmt.Sprintf("%d/%d", ans, len(strSplit))
}
