package main

import "fmt"


func main() {

	cards := newDeck()

	fError := cards.saveToFile("saved-deck")
	fmt.Println(fError)

}
