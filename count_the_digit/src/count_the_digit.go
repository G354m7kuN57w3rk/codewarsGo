package countthedigit

import (
	"strconv"
	"strings"
)

//NbDig square all number between 0 and n. Count number of d
func NbDig(n int, d int) (count int) {
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return count
}
