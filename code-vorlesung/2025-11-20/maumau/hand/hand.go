package hand

import (
	"prog1-material/code-vorlesung/2025-11-20/maumau/card"
	"prog1-material/code-vorlesung/2025-11-20/maumau/deck"
)

type Hand struct {
	Cards []card.Card
}

const (
	HandSize = 7
)

func NewHand(HandSize int, d *deck.Deck) Hand {
	var hand []card.Card
	if len(d.Cards) < HandSize {
		//deck muss neu gemacht werden
	}
	for i := 0; i < HandSize; i++ {
		c, _ := d.Draw()
		hand = append(hand, c)

	}
	return Hand{Cards: hand}
}

func (h Hand) Add(card card.Card) {
	h.Cards = append(h.Cards, card)
}
