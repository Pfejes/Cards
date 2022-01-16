package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (cards deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	randNumber := rand.New(source)

	for i := range cards {
		newPosition := randNumber.Intn(len(cards) - 1)

		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}