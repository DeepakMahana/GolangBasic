package main

func main() {

	// cards := newDeck()
	// hand, remainingCard := deal(cards, 5)
	// hand.print()
	// remainingCard.print()

	// cards := newDeck()
	// cards.saveToFile("My_Cards")

	// cards := newDeckFromFile("My_Cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
