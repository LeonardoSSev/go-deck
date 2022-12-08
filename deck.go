package main

import (
	"fmt"
	"math/rand"
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
