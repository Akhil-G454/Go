package main

import "fmt"

func main() {
	var card = newCard()
	fmt.Println(card)
	var output = customPrint() // Call the renamed function
	fmt.Println(output)
}

func newCard() string {
	return "Ace of spades"
}
