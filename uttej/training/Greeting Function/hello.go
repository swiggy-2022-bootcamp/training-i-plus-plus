package main

import "fmt"

const greetingEnglish = "Hello "
const greetingSpanish = "Hola "
const greetingFrench = "Bonjour "
const greetingTelugu = "Namaskaram "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// if language == "Spanish" {
// 	return greetingSpanish + name
// }
// ==> Refactoring to switch
// if language == "French" {
// 	return greetingFrench + name
// }
func greetingPrefix(language string) (prefix string) {

	switch language {

	case "Spanish":
		prefix = greetingSpanish
	case "French":
		prefix = greetingFrench
	case "Telugu":
		prefix = greetingTelugu
	default:
		prefix = greetingEnglish
	}

	return //no need to specify prefix again as we've already mentioned it while defining the function
}

func main() {

	fmt.Println(Hello("World", ""))

}
