package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
) 

type deck []string
type halfDeck []string

func newDeck () deck {
	cards := deck{}
	cardsSuite := halfDeck{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValue := halfDeck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardValue := range cardsValue {
		for _, cardSuite := range cardsSuite{
			cards = append(cards, cardValue + " Of " + cardSuite)
		}
	}
	return cards
}

func deal (d deck, handSize int) (deck, deck){ //for returning multiple values
	return d[:handSize], d[handSize:] //this works like slice in js
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // when writing to a file you must convert anything to string and that string to a byte slice
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // swapping
	}
}

func (d deck) print() { // is a reciever function as it is having a value (d deck)
	for i, card := range d {
		fmt.Println(i, card)
	}
}