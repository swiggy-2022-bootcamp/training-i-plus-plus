package main

// generics are in beta phase. Install: go install golang.org/dl/go1.18beta1@latest
// type Number interface{
// 	int32 | int64 | float32 | float64
// }

// func betterSum[number Number] (num []number) number{
// 	sum := 0
// 	for _, i := range number{
// 		sum += i
// 	}
// 	return sum
// }

// func sum[number int | float32] (num []number) number{
// 	sum := 0
// 	for _, i := range number{
// 		sum += i
// 	}
// 	return sum
// }

// func main(){
// 	intArray := []int{1,2,3,4,5}
// 	floatArray := []float32{1.2,2,3,4,5.6}

// 	fmt.Println(betterSum(intArray))
// 	fmt.Println(betterSum(floatArray))
// }
