package main

import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) { //function is called automatically by Go test runner
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Ecpected deck length of 16, but got %v", len(d)) //to print value of len(d) inside string put %v
	}

	if d[0] !=  "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in dec, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}