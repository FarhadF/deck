package main

import "fmt"

func main() {
	d := newDeck()
	d = d.shuffle()
	d.print()
	hand, remaining := d.deal(5)
	fmt.Println(len(d))
	fmt.Println(hand, remaining)

}
