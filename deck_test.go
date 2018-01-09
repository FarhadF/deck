package main

import "testing"

func TestDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("expected len 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be ace of spades but got %v", d[0])
	}
	d.shuffle()
	if d[0] == "Ace of Spades" && d[1] == "Two of Spades" && d[2] == "Three of Spades" {
		t.Error("expected deck to shuffle but it seems not, you may need to run the test again cause of bad luck!!")
	}

}
