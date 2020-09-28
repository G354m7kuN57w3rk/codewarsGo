package highestandlowest

import (
	"fmt"
	"strconv"
	"strings"
)

//HighAndLow returns the highest and lowest number of input string
func HighAndLow(in string) string {
	strSplit := strings.Split(in, " ")
	start, _ := strconv.Atoi(strSplit[0])
	min, max := start, start
	for _, v := range strSplit {
		num, _ := strconv.Atoi(v)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}
