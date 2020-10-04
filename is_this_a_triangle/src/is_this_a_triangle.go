package isthisatriangle

//IsTriangle returns true if triangle can be built
func IsTriangle(a, b, c int) bool {
	if a+b > c && c+b > a && a+c > b {
		return true
	}
	return false
}
