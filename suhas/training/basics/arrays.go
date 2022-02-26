package main 


import "fmt"

var numbersArray [3]int

func main() {
	namesArray := [4]string{"John","Joe","Jonathan","Jimmy"}
	
	numbersArray[0]=10
	numbersArray[1]=20
	numbersArray[2]=40
	
	fmt.Println("Numbers Array is :",numbersArray)
	fmt.Println("Names Array is:",namesArray)

	namesArray[1] = "Joey"
	fmt.Println("Names Array is",namesArray)

	nums := [6]int{1,2,3,4,5,6}
	fmt.Println(nums)

	var s []int = nums[1:4] //slice
	fmt.Println(s)
	s[1] = 100
	fmt.Println(s)
	fmt.Println(nums) // nums[2] = 100 
}