package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Changing Variable"

	// card := newCard();
	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// //Iterate through Slices
	// for i, card:= range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)

	// cards := newDeck()
	// cards.print()

	// handDeck, remainingDeck := cards.deal(5)

	// handDeck.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards");

	cards := newDeckFromFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()

	arr := []int {0,1,2,3,4,5,6,7,8,9, 10}

	for _, value := range arr {
		if value % 2 == 0 {
			fmt.Printf("%v is Even", value)
		} else {
			fmt.Printf("%v is Even", value)
		}
		fmt.Println()
	}
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

//returning int from function
// func newCard() int {
// 	return 12
// }