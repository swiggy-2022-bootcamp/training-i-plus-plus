package main

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 10)
	hand.print()
	hand.shuffle()
	hand.print()
}
