package main

func main() {
	cards := newDeck()
	cards, hand := cards.drawHand(5)
	hand.saveToFile("hand.txt")
	new_hand := readDeckFromFile("hand.txt")
	new_hand.print()
}
