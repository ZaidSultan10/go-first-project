package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards but got %v",len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected Ace Of Spades but got %v", d[0])
	}

	if d[len(d) - 1] != "King Of Clubs" {
		t.Errorf("Expected King Of Clubs but got %v", d[len(d) - 1])
	}
}

func TestNewForSaveToFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("File test : Expected 52 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}