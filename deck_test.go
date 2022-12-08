package main

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := CreateDeck()

	if len(d) != 52 {
		t.Errorf("The deck should have %v cards, got %v cards", 52, len(d))
	}

	// Expected first card
	efc := "Ace of Spades"

	// First deck's card
	fdc := d[0]

	if fdc != efc {
		t.Errorf("The first card of the deck should be %v, got %v", efc, fdc)
	}

	// Expected last card
	elc := "King of Diamonds"

	// Last deck's card
	ldc := d[len(d)-1]

	if ldc != elc {
		t.Errorf("The last card of the deck should be %v, got %v", elc, ldc)
	}
}
