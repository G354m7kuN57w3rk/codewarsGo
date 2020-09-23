package grasshoppersummation

//Summation sum every number from 1 to num
func Summation(n int) int {
	var ans int = 0
	for i := 0; i < n+1; i++ {
		ans += i
	}
	return ans
}
