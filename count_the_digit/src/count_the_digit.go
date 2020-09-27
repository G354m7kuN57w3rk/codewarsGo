package countthedigit

import (
	"strconv"
	"strings"
)

//NbDig square all numbers between 0 and n and count the number of d
func NbDig(n int, d int) (count int) {
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return count
}
