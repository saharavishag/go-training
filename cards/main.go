package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards)
	cards.print()

	// split the deck
	hand, remaningDeck := deal(cards, 5)

	fmt.Println()
	fmt.Println("hand:")
	hand.print()
	fmt.Println()
	fmt.Println("remaningDeck:")
	remaningDeck.print()

	// type convertion
	greeting := "hi there"
	fmt.Println([]byte(greeting))

	// print tostring
	fmt.Println(cards.toString())

	// save this deck
	cards.saveToFile("my_cards")

	// create from file
	fmt.Println("Deck from file...")
	newDeckFromFile("my_cards").print()

	// shuffle
	cards.shuffle()
	cards.print()
}
