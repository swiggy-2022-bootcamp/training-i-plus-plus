package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {} 
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}


//if receivers are not used inside the function we can also remove it to avaoid "Not used" erros
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

//With interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//Without Interfaces

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//Other Example
// type bot interface {
// 	getGreeting(string, int) (string, error)
// 	getBotVersion() float64
// 	respondToUser(user) string
// }