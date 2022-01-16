package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// no receiver, because it creates the set
func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

// deal handSize pieces from the whole deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(cards.toString()), 0666)
}
