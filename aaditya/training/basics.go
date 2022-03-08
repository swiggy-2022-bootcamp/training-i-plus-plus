package lasagna

import "fmt"

//Concepts - Packages, variables, constans, functions

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven ;
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers*2;
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return numberOfLayers*2 + actualMinutesInOven;
}

func main(){
	rt := RemainingOvenTime(12)
	fmt.Println("Remaining time to cook lasagna ", rt)
	pt := PreparationTime(3)
	fmt.Println("Preparation time to cook lasagna ", pt)
	et := ElapsedTime(3,12)
	fmt.Println("Elapsed time to cook lasagna ", et)
}