package main

import "fmt"

func main() {
	cards := newDeck()
	cards, hand := cards.drawHand(5)
	cards.print()
	hand.print()
	fmt.Println(hand.toString())
}
