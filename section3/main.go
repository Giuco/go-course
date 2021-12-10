package main

func main() {
	cards := newDeck()
	cards, hand := cards.drawHand(5)
	hand.saveToFile("hand.txt")
}
