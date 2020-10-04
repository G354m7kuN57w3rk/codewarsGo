package makeafunctionthatdoesarithmetic

//Arithmetic returns the result of number a and b depending of operator
func Arithmetic(a int, b int, operator string) int {
	var ans int
	switch {
	case operator == "add":
		ans = a + b
	case operator == "subtract":
		ans = a - b
	case operator == "multiply":
		ans = a * b
	case operator == "divide":
		ans = a / b
	}
	return ans
}
