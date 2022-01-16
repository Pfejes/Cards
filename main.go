package main

import "fmt"

func main() {

	// slice
	cards := []string{"a", "b", newCard()}
	cards = append(cards, "new element")

	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}