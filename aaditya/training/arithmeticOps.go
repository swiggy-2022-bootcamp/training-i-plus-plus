package cars

//Concepts - Arithemtic operation, different number types in golang


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var n float64 = successRate/100 * float64(productionRate)
	return n
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var n float64 = float64(productionRate)/60
    n = n * successRate/100
    return int(n)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var groups int = carsCount/10;
    var individuals int = carsCount % 10;
    var cost uint = uint(groups) * 95000 + uint(individuals) * 10000
    return cost
}