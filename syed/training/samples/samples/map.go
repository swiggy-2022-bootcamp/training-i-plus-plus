package main
import "fmt"

var nameAgeMap map[string]int

func main() {
	nameAgeMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}
	fmt.Println("Print the age of James: ", nameAgeMap["James"])

	//We can range through the map and print each value:
	for key, value := range nameAgeMap {
		fmt.Printf("%v is %d years old\n", key, value)
	}
}