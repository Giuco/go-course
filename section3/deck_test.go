package main

import "testing"

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
