package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of " +suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card :=  range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(([]string(d)), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) //0666 - anyone can read or write the file
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit((1)) //error code of 1
	} 

	s := strings.Split(string(bs), ",")
	return deck(s)  
}

func (d deck) shuffle() {
	
	source := rand.NewSource(time.Now().UnixNano())
	randomSeed := rand.New(source)
	 
	for index := range d {
		newPosition := randomSeed.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}