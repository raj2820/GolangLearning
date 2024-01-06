package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()} //slice
	cards = append(cards, "six of spades")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}