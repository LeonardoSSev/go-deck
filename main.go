package main

import "fmt"

func main() {
	d := CreateDeck()
	fmt.Println("Creating a deck")
	d.Print()

	fmt.Println("Saving my deck")
	SaveDeck("_mydeck.txt", d)

	fmt.Println("Importing the deck from a file")
	d2 := ImportDeck("_mydeck.txt")

	fmt.Println("Shuffling the deck")
	d2.Shuffle()
	d2.Print()

	fmt.Println("Dealing the cards")
	h, rd := Deal(d2, 5)

	fmt.Println("My hands")
	h.Print()

	fmt.Println("The rest of the deck")
	rd.Print()
}
