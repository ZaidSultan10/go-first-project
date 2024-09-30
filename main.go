package main

import "fmt"

func main(){
	cards := readFromFile("my_cards")
	// cards.saveToFile("my_cards")
	cards.shuffleDeck()
	cards.print()
	// fmt.Println(cards.toString())
	// firstDeck, secondDeck := deal(cards, 5)
	// firstDeck.print()
	// fmt.Printf("\n ============= \n")
	// secondDeck.print()
}

func generateCard(card string) string {
	return card
}

func sum(a int, b int) int {
	var c int = a + b
	return c
}

func randomFunc() {
	fmt.Println("apples")
}