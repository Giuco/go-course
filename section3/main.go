package main

func main() {
	cards := newDeck()
	cards, hand := cards.drawHand(5)
	cards.print()
	hand.print()
}
