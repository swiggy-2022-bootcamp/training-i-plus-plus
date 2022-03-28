package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Create and return a deck
func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Jack", "Queen"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print all cards in the deck
func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

// Deal will split the deck into 2 slices of length n and len(deck) - n
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert deck to a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save deck to a file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// New Deck from file
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// Shuffle the deck in random order
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < len(d); i++ {
		var newPosition = r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
