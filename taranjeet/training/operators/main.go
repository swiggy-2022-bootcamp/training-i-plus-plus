package operators

func Operators(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	prod := a * b
	div := a / b
	return sum, sub, prod, div
}

func ConditionalOperators(a, b int) (bool, bool, bool, bool) {

	return a > b, a == b, a > 0 && a != b, a > 0 || a == b

}

func Calculator(symbol string, a, b int) int {
	switch symbol {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0

	}
}
