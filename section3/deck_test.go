package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected_len := 52

	if len(d) != expected_len {
		t.Errorf("Expected deck size of %v but got %v", expected_len, len(d))
	}
}
