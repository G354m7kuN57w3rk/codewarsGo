package maximumlengthdifference

import "math"

//MxDifLg returns max length difference between strings of a1 and a2
func MxDifLg(a1 []string, a2 []string) int {
	ans := -1
	for _, x := range a1 {
		for _, y := range a2 {
			if v := math.Abs(float64(len(x)) - float64(len(y))); v > float64(ans) {
				ans = int(v)
			}
		}
	}
	return ans
}
