package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := newDeck()
	cards.shuffle()
	_, hand := drawHand(cards, 5)
	hand.saveToFile("hand.txt")
	loadedHand := readDeckFromFile("hand.txt")
	loadedHand.print()
}
