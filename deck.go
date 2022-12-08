package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func CreateDeck() deck {
	// Deck, Suits and Cards
	d := []string{}
	ss := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cs := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range ss {
		for _, c := range cs {
			d = append(d, c+" of "+s)
		}
	}

	return d
}

func (d deck) Print() {
	for i, cs := range d {
		fmt.Printf("%v - %v\n", i, cs)
	}
}

func (d deck) Shuffle() {
	for i, _ := range d {
		// Generating a random generator from a dynamic seed
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		// Getting a random card (rc) from the deck with a random index (ri)
		ri := r.Intn(len(d) - 1)
		rc := d[ri]

		// Shuffling the cards
		d[i], d[ri] = rc, d[i]
	}
}

func SaveDeck(filename string, d deck) {
	das := getDeckAsString(d)

	err := os.WriteFile(filename, []byte(das), 0666)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
}

func getDeckAsString(d deck) string {
	return strings.Join(d, ",")
}

func ImportDeck(filename string) deck {
	dab, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	return getDeckFromString(string(dab))
}

func getDeckFromString(das string) deck {
	return strings.Split(das, ",")
}

func Deal(d deck, hs int) (deck, deck) {
	return d[:hs], d[hs:]
}
