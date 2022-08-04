package main

import "fmt"

func main() {
	// var card string =  "Ace of Spades"
	cards := newDeck()
	fmt.Println("Create a new cards ", cards)

	// test deal method return 2 deck instance 
	// --> can call print() method
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	greeting := "Hi there!"
	// slice of greeting which is converted to byte
	fmt.Println("Convert to Byte ", []byte(greeting))

	// test save method
	cards.saveToFile("my_cards")

	// load new cards from a file
	new_cards := newDeckFromFile("my_cards")
	new_cards.shuffle()
	new_cards.print()
}