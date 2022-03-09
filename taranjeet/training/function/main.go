package function

func Fib(size int) (finalFib []int) {
	//var fib []int
	//fib = append(fib, 0)
	//fib = append(fib, 1)
	fib := make([]int, size)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < size; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib

}

func TestFunction(str1, str2 string, num1 int) (finalString string, secString string) {
	for i := 0; i < num1; i++ {
		str1 = str1 + str2
	}
	return str1, str2
}
