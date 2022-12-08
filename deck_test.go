package main

import (
	"os"
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

func TestShuffle(t *testing.T) {
	d := CreateDeck()
	d2 := CreateDeck()

	d.Shuffle()
	d2.Shuffle()

	das := getDeckAsString(d)
	das2 := getDeckAsString(d2)

	if das == das2 {
		t.Error("The decks should not be equal")
	}
}

func TestSaveDeck(t *testing.T) {
	filename := "mydeck.txt"
	_ = os.Remove(filename)

	d := CreateDeck()
	SaveDeck(filename, d)

	_, err := os.ReadFile(filename)

	if err != nil {
		t.Error("Deck not saved")
	}
}

func TestImportDeck(t *testing.T) {
	filename := "mydeck.txt"
	_ = os.Remove(filename)

	d := CreateDeck()
	SaveDeck(filename, d)

	d2 := ImportDeck(filename)

	das := getDeckAsString(d)
	das2 := getDeckAsString(d2)

	if das != das2 {
		t.Error("Deck not imported properly")
	}
}
