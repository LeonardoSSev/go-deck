package main

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
