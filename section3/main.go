package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards, hand := cards.drawHand(5)
	hand.saveToFile("hand.txt")
	new_hand := readDeckFromFile("hand.txt")
	new_hand.print()
}
