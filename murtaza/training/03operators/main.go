package operators

import "math"

//-------- Mathematical Operators in Go
func Add(a int, b int) int {
	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

func modulo(a, b float64) float64 {
	return math.Mod(a, b)
}

//-------- Logical Operators in Go
func IsCreditScoreEligible(creditScore float32) bool {
	return creditScore >= 700.0
}

func IsCouponApplicable(couponCode string, isNewUser bool) bool {
	return couponCode == "NEWUSER50" && isNewUser
}

func GetDeliveryFee(distanceInKm float32, timeTakenInMinutes float32) float32 {
	if distanceInKm > 5 || timeTakenInMinutes > 15 {
		return 15.0
	}
	return 0.0
}

//------ Conditional operator
func calculator(operatorSymbol string, a int, b int) float32 {
	switch operatorSymbol {
	case "+":
		return float32(Add(a, b))
	case "-":
		return float32(Substract(a, b))
	case "*":
		return float32(Multiply(a, b))
	case "/":
		return float32(Divide(a, b))
	case "%":
		return float32(modulo(float64(a), float64(b)))
	default:
		return 0.0
	}
}
