package main

import "fmt"

func main() {

	// slice
	cards := []string{"a", "b", newCard()}
	cards = append(cards, "new element")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}