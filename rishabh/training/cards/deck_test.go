package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be \"Ace of spades\", but got %v", d[0])
	}

	if d[len(d)-1] != "Queen of Clubs" {
		t.Errorf("Expected last card to be \"Queen of clubs\", but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testingFileName := "_duckTesting"
	os.Remove(testingFileName)

	d := newDeck()
	d.saveToFile(testingFileName)

	loadedDeck := newDeckFromFile(testingFileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(loadedDeck))
	}

	os.Remove(testingFileName)
}
