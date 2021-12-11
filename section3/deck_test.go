package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected_len := 52

	if len(d) != expected_len {
		t.Errorf("Expected deck size of %v but got %v", expected_len, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', go %v", d[len(d)-1])
	}
}

func TestSaveAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(deck) != len(loadedDeck) {
		t.Errorf("Both sizes should be equal. deck=%v. loadedDeck=%v", len(deck), len(loadedDeck))
	}

	for i := range deck {
		if deck[i] != loadedDeck[i] {
			t.Errorf("Deck are diffent at index %v. deck=%v. loadedDeck=%v", i, len(deck), len(loadedDeck))
		}
	}

	os.Remove("_decktesting")
}
