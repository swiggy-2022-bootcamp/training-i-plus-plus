package sonarqube

import (
	"fmt"
	"sonarqube/logTest"
	"testing"
)

func main() {
	fmt.Println("Started Test")
	if logTest.Test(&testing.T{}) == true {
		fmt.Println("Test Passed")
	} else {
		fmt.Println("Test Failed")
	}
}
