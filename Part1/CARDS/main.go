package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" 
	// card :=  "Ace of Spades"  //short way of declarition variable, used only when defining a new value and not assignment
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}