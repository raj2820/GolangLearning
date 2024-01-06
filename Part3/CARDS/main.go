package main

// import "fmt"

func main() {
	cards := newDeck()

	hand, remaoningsCard := deal(cards, 5)
	hand.print()
	remaoningsCard.print()
	// fmt.Println(cards)
}

