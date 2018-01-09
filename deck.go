package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	var d deck
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Soldier", "Queen", "King"}
	for _, s := range suits {
		for _, v := range values {
			d = append(d, v+" of "+s)
		}

	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
