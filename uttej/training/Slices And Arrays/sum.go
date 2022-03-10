package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { /* "_" is a blank identifier. range gives an index and value at index.
		We are choosing to ignore the index by using _ */
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	// lengthOfNumbers := len(numbersToSum)

	// sums1 := make([]int, lengthOfNumbers)

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSumTails ...[]int) []int {

	var sumTails []int

	for _, numbers := range numbersToSumTails {
		if len(numbers) == 0 {
			sumTails = append(sumTails, 0)
		} else {
			tail := numbers[1:]
			sumTails = append(sumTails, Sum(tail))
		}
	}
	return sumTails
}
