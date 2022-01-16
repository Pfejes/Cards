package main


func main() {

	// slice
	cards := newDeck()
	
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
