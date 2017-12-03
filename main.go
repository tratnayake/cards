//Starting Section 3 Lecture 16
package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	//------
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//-----
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	//-----
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
