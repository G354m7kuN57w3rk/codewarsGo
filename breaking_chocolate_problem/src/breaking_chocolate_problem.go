package breakingchocolateproblem

//BreakChocolate returns number of breaks needed to split a chocolate bar n x m
func BreakChocolate(n, m int) int {
	ans := n*m - 1
	if ans < 0 {
		return 0
	}
	return ans
}
