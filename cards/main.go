package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// number := newInt()
	// card = "Five of Diamonds"
	// fmt.Println(card, number)

	// a slice of string, slice是go的一种数组结构
	// greeting := "Hi!"
	// fmt.Println([]byte(greeting))
	cards := newDeckFromFile("my_cards")
	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	// hand.print()
	// remainingCards.print()
}

func newInt() int {
	return 12
}
